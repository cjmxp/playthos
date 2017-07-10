// +build desktop,audio

package audio

import (
	"github.com/autovelop/playthos"
	"golang.org/x/mobile/exp/audio/al"
	"io"
	"log"
	"time"
)

func init() {
	engine.NewSystem(&Audio{})
}

type Audio struct {
	engine.System
	buffers []al.Buffer
	sources []al.Source
}

func (a *Audio) InitSystem() {
	al.OpenDevice()
	al.SetListenerPosition([3]float32{0, 0, 0})
}

func NewSound() *Sound {
	s := &Sound{}
	s.SetActive(true)
	return s
}

func (a *Audio) PlaySound(sound *Sound) {
	if sound.Active() && sound.ready && !sound.playing {
		sound.playing = true

		source := a.sources[sound.idx]

		al.PlaySources(source)

		if sound.Loops() {
			source.Seti(paramLooping, 1)
		} else {
			go func(sound *Sound) {
				time.Sleep(sound.audioFile.Duration())
				sound.playing = false
			}(sound)
		}
	}
}

func (a *Audio) NewComponent(sound engine.ComponentRoutine) {
	switch sound := sound.(type) {
	case *Sound:
		source := al.GenSources(1)[0]
		a.sources = append(a.sources, source)
		if code := al.Error(); code != 0 {
			log.Fatalf("audio: cannot generate an audio source [err=%x]", code)
		}

		buffer := al.GenBuffers(1)[0]
		a.buffers = append(a.buffers, buffer)
		audioFile := sound.Get()
		buf := make([]byte, audioFile.FileSize())
		size := int64(0)
		for {
			n, err := audioFile.GetReader().Read(buf)
			if n > 0 {
				size += int64(n)
				buffer.BufferData(al.FormatStereo16, buf[:n], int32(audioFile.SampleRate()))
			}
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
		}

		source.QueueBuffers(buffer)

		sound.audioSystem = a
		sound.idx = len(a.sources) - 1
		sound.ready = true
		if sound.playOnReady {
			a.PlaySound(sound)
		}
		break
	}
}

func (a *Audio) ComponentTypes() []engine.ComponentRoutine {
	return []engine.ComponentRoutine{&Sound{}}
}

func (a *Audio) NewIntegrant(integrant engine.IntegrantRoutine) {}

/*
The golang experimental openal bindings for have looping as a const yet. So I just guessed it based on the below C code.
https://github.com/kcat/openal-soft/blob/0a361fa9e27b9d9533dffe34663efc3669205b86/Alc/ALc.c
*/
const (
	paramLooping = 0x1007
)