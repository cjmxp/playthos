package editor

import (
	"gde/engine"
	// "gde/network"
	"encoding/json"
	"gde/render"

	// "gde/render/animation"
	"fmt"
	"gde/render/ui"
	"github.com/gorilla/websocket"
	"log"
	"net"
	"net/http"
	"strconv"
)

type Scene struct {
	name string
}

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
	return true
}}

type EditorAction struct {
	Action uint
	Data   string
}

type EditorUpdate struct {
	Entity      string
	Component   string
	Property    string
	SubProperty string
	Value       string
}

func (s *Scene) CreateEditorServer(game *engine.Engine) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		// defer c.Close()
		for {
			mt, message, err := ws.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}

			var editor_action EditorAction
			err_unmarshal := json.Unmarshal(message, &editor_action)
			if err_unmarshal != nil {
				fmt.Println("error:", err_unmarshal)
				fmt.Println("error:", string(message))
			}

			switch editor_action.Action {
			case 0:
				entity_json, err := json.Marshal(game.GetEntity("Player"))
				if err != nil {
					fmt.Println("error:", err)
				}

				err = ws.WriteMessage(mt, entity_json)
				if err != nil {
					log.Println("write:", err)
					break
				}
			case 1:
				fmt.Printf("%v", editor_action.Data)
				var editor_update EditorUpdate
				err_update_unmarshal := json.Unmarshal([]byte(editor_action.Data), &editor_update)
				if err_update_unmarshal != nil {
					fmt.Println("error:", err)
				}
				// fmt.Printf("%+v", game.GetEntity(editor_update.Entity))
				fmt.Printf("%+v", game.GetEntity(editor_update.Entity).GetComponentByStr(editor_update.Component).GetProperty(editor_update.Property))
				switch editor_update.Component {
				case "*render.Transform":
					{
						vec3 := game.GetEntity(editor_update.Entity).GetComponentByStr(editor_update.Component).GetProperty(editor_update.Property)
						switch vec3 := vec3.(type) {
						case render.Vector3:
							switch editor_update.SubProperty {
							case "X":
								v, err := strconv.ParseFloat(editor_update.Value, 64)
								if err != nil {
									fmt.Println("error:", err)
								}
								vec3.X = float32(v)
								break
							case "Y":
								v, err := strconv.ParseFloat(editor_update.Value, 64)
								if err != nil {
									fmt.Println("error:", err)
								}
								vec3.Y = float32(v)
								break
							case "Z":
								v, err := strconv.ParseFloat(editor_update.Value, 64)
								if err != nil {
									fmt.Println("error:", err)
								}
								vec3.Z = float32(v)
								break
							}
							game.GetEntity(editor_update.Entity).GetComponentByStr(editor_update.Component).SetProperty(editor_update.Property, vec3)
						}
					}
				}

				// game.GetEntity(editor_update.Entity).GetComponentByStr(editor_update.Component).GetProperty(
				// var vec3 render.Vector3
				// err_vec3_unmarshal := json.Unmarshal([]byte(editor_update.Value), &vec3)
				// if err_vec3_unmarshal != nil {
				// 	fmt.Println("error:", err_vec3_unmarshal)
				// }
			}

			// log.Printf("recv: %s", message)
			// err = ws.WriteMessage(mt, message)
			// if err != nil {
			// 	log.Println("write:", err)
			// 	break
			// }
		}

		// GET ON CONNECT EVENT
		// SEND JSON OF ENTITIES TO CLIENT
		// ALLOW CLIENTS TO SEND NEW VALUES OF PROPERTIES BY PROPERTY NAME AND ENTITY ID AS JSON VALUE (SHOULD BE CONVERTED ACCORDINGLY)
		// BROADCAST UPDATE TO ALL CLIENTS

		// ticker := time.NewTicker(time.Second)
		// defer ticker.Stop()
		// for {
		// 	select {
		// 	case t := <-ticker.C:
		// 		err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
		// 		if err != nil {
		// 			log.Println("write:", err)
		// 			return
		// 		}
		// 	}
		// }
	})
	var ip net.IP
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Println(err)
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			log.Println(err)
		}
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
		}
	}
	address := "192.168.43.1:8080"
	if len(fmt.Sprintf("%v", ip)) < 25 {
		address = fmt.Sprintf("%v:8080", ip)
	}
	go func(address string) {
		fmt.Println(address)
		log.Fatal(http.ListenAndServe(address, nil))
	}(address)

	// THIS IS NOT REALLY REQUIRED BUT JUST USED TO DEBUGGING ON LOCAL MACHINE (testing broadcast etc.)
	// network := &network.Network{ServerIP: address}
	// game.AddSystem(engine.SystemNetwork, network)
	// network.Init()
}

func (s *Scene) LoadScene(game *engine.Engine) {
	sys_render, err := game.GetSystem(engine.SystemRender).(render.RenderRoutine)
	if !err {
		log.Printf("\n\n ### ERROR ### \n%v\n\n", err)
		return
	}

	// Simple Quad mesh renderer
	comp_renderer := &render.MeshRenderer{}
	comp_renderer.Init()
	comp_renderer.LoadMesh(&render.Mesh{
		Vertices: []float32{
			0.2, 0.2, 0.0, 1.0, 0.0, 0.0, 1.0, 1.0,
			0.2, 0.0, 0.0, 0.0, 1.0, 0.0, 1.0, 0.0,
			0.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0,
			0.0, 0.2, 0.0, 0.0, 1.0, 1.0, 0.0, 1.0,
		},
		Indicies: []uint8{
			0, 1, 3,
			1, 2, 3,
		},
	})
	texture := &render.Texture{}
	texture.NewTextureMobile("weapon.png")
	comp_renderer.LoadTexture(texture)

	sys_render.LoadRenderer(comp_renderer)

	// Create player entity
	ent_player := &engine.Entity{Id: "Player"}
	ent_player.Init()
	ent_player.Add(game)

	ent_player_comp_transform := &render.Transform{}
	ent_player_comp_transform.Init()
	ent_player_comp_transform.SetProperty("Position", render.Vector3{0.5, 1.0, 1.0})
	ent_player_comp_transform.SetProperty("Rotation", render.Vector3{0, 0, 45})

	ent_player.AddComponent(ent_player_comp_transform)
	ent_player.AddComponent(comp_renderer)

	// Create UI entity

	// // First create a UI system
	sys_render.AddUISystem(game)
	sys_ui, err := game.GetSystem(engine.SystemUI).(ui.UIRoutine)
	if !err {
		log.Printf("\n\n ### ERROR ### \n%v\n\n", err)
		return
	}
	log.Printf("UI SYSTEM: %+v", sys_ui)

	// Load a simple font
	font := &ui.Font{}
	font.NewFont()

	ent_box := &engine.Entity{Id: "Box"}
	ent_box.Init()
	ent_box.Add(game)

	ent_box_comp_transform := &render.Transform{}
	ent_box_comp_transform.Init()
	ent_box_comp_transform.SetProperty("Position", render.Vector3{120, 200, 0.5})
	ent_box_comp_transform.SetProperty("Dimensions", render.Vector2{240, 400})
	ent_box.AddComponent(ent_box_comp_transform)

	comp_ui_renderer := &ui.UIRenderer{}
	comp_ui_renderer.Init()

	text := &ui.Text{}
	text.SetFont(font)
	text.SetText(`Common Sword`)
	comp_ui_renderer.SetProperty("Text", text.TextToVec4())
	comp_ui_renderer.SetProperty("Scale", 2.0)
	// comp_ui_renderer.SetProperty("Padding", render.Vector4{0 /*top*/, 0 /*right*/, 0 /*bottom*/, 0 /*left*/})

	ent_box.AddComponent(comp_ui_renderer)

	sys_ui.LoadRenderer(comp_ui_renderer)

	s.CreateEditorServer(game)

	// Officially the worst Animation system since forever!
	// sys_anim := &animation.Animation{}
	// sys_anim.Init()
	// game.AddSystem(engine.SystemAnimation, sys_anim)

	// ent_box_comp_animator := &animation.Animator{EndFrame: 240, Start: func(frame int) {
	// 	ent_box_comp_transform.SetProperty("Dimensions", render.Vector2{80, 600})
	// }, Step: func(frame int) {
	// 	dimensions := ent_box_comp_transform.GetProperty("Dimensions")
	// 	switch dimensions := dimensions.(type) {
	// 	case render.Vector2:
	// 		ent_box_comp_transform.SetProperty("Dimensions", render.Vector2{dimensions.X + 1, dimensions.Y})
	// 	}
	// }}
	// ent_box_comp_animator.Init()
	// ent_box.AddComponent(ent_box_comp_animator)

	// // // Lets test keyboard support
	// keyInput, err := engine.GetSystem(SystemInputKeyboard).(Input)
	// if !err {
	// 	log.Println(err)
	// 	return
	// }

	// // Right arrow
	// keyInput.BindOn(262, func() {
	// 	box2_position.X += 0.1
	// 	box2_transform.SetProperty("Position", box2_position)
	// 	textTmp := &Text{text.GetText() + "Right", font}
	// 	text_renderer.SetProperty("Text", textTmp.TextToVec2())
	// })

	// // Left arrow
	// keyInput.BindOn(263, func() {
	// 	box2_position.X -= 0.1
	// 	box2_transform.SetProperty("Position", box2_position)
	// })

	// // Up arrow
	// keyInput.BindOn(265, func() {
	// 	box2_position.Y -= 0.1
	// 	box2_transform.SetProperty("Position", box2_position)
	// })

	// // Down arrow
	// keyInput.BindOn(264, func() {
	// 	box2_position.Y += 0.1
	// 	box2_transform.SetProperty("Position", box2_position)
	// })

	// // Lets test pointer support
	// pointerInput, err := engine.GetSystem(SystemInputPointer).(Input)
	// if !err {
	// 	log.Printf("Pointer Input system not started/found\nERROR:\n%v\n\n", err)
	// 	return
	// }

	// // pointer Move
	// pointerInput.BindMove(func(x float64, y float64) {
	// 	box2_position.X = float32(x/360) * 1
	// 	box2_position.Y = float32(y/640) * 2
	// 	box2_transform.SetProperty("Position", box2_position)
	// })

	// Left click
	// pointerInput.BindAt(0, func(x float64, y float64) {
	// 	box2_position.Y += 0.1
	// 	box2_transform.SetProperty("Position", box2_position)
	// })

	// // Right click
	// pointerInput.BindAt(1, func(x float64, y float64) {
	// 	box2_position.Y += 0.1
	// 	box2_transform.SetProperty("Position", box2_position)
	// })

	// // Lets test touch support
	// touchInput, err := engine.GetSystem(SystemInputTouch).(InputRoutine)
	// if !err {
	// 	log.Printf("Touch Input system not started/found\nERROR:\n%v\n\n", err)
	// 	return
	// }

	// // Touch down
	// touchInput.BindAt(0, func(x float64, y float64) {
	// 	box2_position.Y += 0.1
	// 	box2_transform.SetProperty("Position", box2_position)
	// })

	// // Touch up
	// touchInput.BindAt(1, func(x float64, y float64) {
	// 	box2_position.Y += 0.1
	// 	box2_transform.SetProperty("Position", box2_position)
	// })
}
