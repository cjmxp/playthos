// +build render

package render

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/std"
)

type Render struct {
	engine.System
	// RenderRoutine
}

func NewRenderSystem(render RenderRoutine) {
	engine.NewSystem(render)
}

type RenderRoutine interface {
	// engine.SystemRoutine
	engine.Updater
	NewShader(vs string, fs string) uint32
	// UnRegisterEntity(*engine.Entity)
	RegisterTransform(*std.Transform)
	RegisterMesh(*Mesh)
	RegisterMaterial(*Material)
	RegisterCamera(*Camera)
}

const (
	VSHADER = `#version 450 core
	layout (location = 0) in vec4 pos;
	layout (location = 1) in vec3 col;
	layout (location = 2) in vec2 tex;

	uniform mat4 model;
	uniform mat4 view;
	uniform mat4 projection;

	// try use in/out later
	layout (location = 0) out vec3 colOut;

	// varying bool hasTexOut;
	layout (location = 1) out vec2 texOut;

	void main( void ) {
		gl_Position = projection * view * model * pos;
		// gl_Position = vec4(1.0, 1.0, 0.0, 1.0);
		colOut = col;
		// hasTexOut = hasTex;
		texOut = tex;
	}
	` + "\x00"

	FSHADER = `#version 450 core

	uniform vec4 color;
	layout (location = 0) in vec3 colOut;

	uniform int hasTexture;
	uniform sampler2D textu;
	layout (location = 1) in vec2 texOut;
	uniform vec2 texOff;

	layout (location = 0) out vec4 fragColor;

	void main() {
		if (hasTexture == 1) {
			vec2 p = texOut + texOff;
			vec4 frag_texture = texture(textu, p) * color;
			if(frag_texture.a < 0.1) {
				discard;
			}
			fragColor = frag_texture;
		} else {
			vec4 frag_color = vec4(colOut, 1.0) * color;
			fragColor = frag_color;
		}
	}
	` + "\x00"
	VSHADER41 = `#version 410 core
	layout (location = 0) in vec4 pos;
	layout (location = 1) in vec3 col;
	layout (location = 2) in vec2 tex;

	uniform mat4 model;
	uniform mat4 view;
	uniform mat4 projection;

	// try use in/out later
	layout (location = 0) out vec3 colOut;

	// varying bool hasTexOut;
	layout (location = 1) out vec2 texOut;

	void main( void ) {
		gl_Position = projection * view * model * pos;
		// gl_Position = vec4(1.0, 1.0, 0.0, 1.0);
		colOut = col;
		// hasTexOut = hasTex;
		texOut = tex;
	}
	` + "\x00"

	FSHADER41 = `#version 410 core

	uniform vec4 color;
	layout (location = 0) in vec3 colOut;

	uniform int hasTexture;
	uniform sampler2D textu;
	layout (location = 1) in vec2 texOut;

	layout (location = 0) out vec4 fragColor;

	void main() {
		if (hasTexture == 1) {
			vec4 frag_texture = texture(textu, texOut) * color;
			if(frag_texture.a < 0.1) {
				discard;
			}
			fragColor = frag_texture;
		} else {
			vec4 frag_color = vec4(colOut, 1.0) * color;
			fragColor = frag_color;
		}
	}
	` + "\x00"
	VSHADER33 = `#version 330 core
	#extension GL_ARB_separate_shader_objects : enable

	layout (location = 0) in vec4 pos;
	layout (location = 1) in vec3 col;
	layout (location = 2) in vec2 tex;

	uniform mat4 model;
	uniform mat4 view;
	uniform mat4 projection;

	// try use in/out later
	layout (location = 0) out vec3 colOut;

	// varying bool hasTexOut;
	layout (location = 1) out vec2 texOut;

	void main( void ) {
		gl_Position = projection * view * model * pos;
		// gl_Position = vec4(1.0, 1.0, 0.0, 1.0);
		colOut = col;
		// hasTexOut = hasTex;
		texOut = tex;
	}
	` + "\x00"

	FSHADER33 = `#version 330 core
	#extension GL_ARB_separate_shader_objects : enable

	uniform vec4 color;
	layout (location = 0) in vec3 colOut;

	uniform int hasTexture;
	uniform sampler2D textu;
	layout (location = 1) in vec2 texOut;

	layout (location = 0) out vec4 fragColor;

	void main() {
		if (hasTexture == 1) {
			vec4 frag_texture = texture(textu, texOut) * color;
			if(frag_texture.a < 0.1) {
				discard;
			}
			fragColor = frag_texture;
		} else {
			vec4 frag_color = vec4(colOut, 1.0) * color;
			fragColor = frag_color;
		}
	}
	` + "\x00"
)
