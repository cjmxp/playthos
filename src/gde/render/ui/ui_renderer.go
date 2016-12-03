package ui

import (
	"gde/engine"
	"gde/render"
	"image"
	"log"
)

type UIRenderer struct {
	engine.Component
	render.RendererRoutine
	Mesh    *render.Mesh
	Texture *render.Texture
}

func (r *UIRenderer) Init() {
	log.Printf("UIRenderer > Init")
	r.Properties = make(map[string]interface{})
	// r.Mesh = &render.Mesh{
	// 	Vertices: []float32{
	// 		1.0, 2.0, -0.1, 1.0, 0.0, 0.0,
	// 		1.0, 0.0, -0.1, 0.0, 1.0, 0.0,
	// 		0.0, 0.0, -0.1, 0.0, 0.0, 1.0,
	// 		0.0, 2.0, -0.1, 0.0, 1.0, 1.0,
	// 	},
	// 	Indicies: []uint8{
	// 		0, 1, 3,
	// 		1, 2, 3,
	// 	},
	// }
}

func (r *UIRenderer) GetProperty(key string) interface{} {
	log.Printf("UIRenderer > Property > Get: %v", key)
	return r.Properties[key]
}

func (r *UIRenderer) SetProperty(key string, val interface{}) {
	log.Printf("UIRenderer > Property > Set: %v", key)
	r.Properties[key] = val
}

func (r *UIRenderer) LoadMesh(mesh *render.Mesh) {
	log.Printf("UIRenderer > Mesh > Load: %v", mesh)
	r.Mesh = mesh
}

func (r *UIRenderer) LoadTexture(texture *render.Texture) {
	log.Printf("UIRenderer > Texture > Load: %v", texture)
	r.Texture = texture
}

// Make this happen on the Render System
func (r *UIRenderer) MeshVertices() []float32 {
	log.Printf("UIRenderer > Mesh > Vectices: %v", len(r.Mesh.Vertices))
	return r.Mesh.Vertices
}

func (r *UIRenderer) MeshIndicies() []uint8 {
	log.Printf("UIRenderer > Mesh > Indicies: %v", len(r.Mesh.Indicies))
	return r.Mesh.Indicies
}

func (r *UIRenderer) TextureRGBA() *image.RGBA {
	log.Printf("UIRenderer > Texture > RGBA: %v", r.Texture.RGBA)
	return r.Texture.RGBA
}