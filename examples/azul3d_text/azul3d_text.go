package main

import (
	"code.google.com/p/azul3d/chippy"
	"code.google.com/p/azul3d/chippy/keyboard"
	"code.google.com/p/azul3d/engine"
	"code.google.com/p/azul3d/event"
	"code.google.com/p/azul3d/math"
	"code.google.com/p/azul3d/scene"
	"code.google.com/p/azul3d/scene/text"
	"image/color"
	_ "image/png"
	"log"
	"runtime"
)

var (
	textNode = scene.New("textNode")
)

// Event handler to manipulate cubes when cursor is moved
func onCursorPosition(ev *event.Event) {
	pos := ev.Data.(*chippy.CursorPositionEvent)

	// If the cursor is not grabbed, we do not transform cubes.
	if !engine.Window.CursorGrabbed() {
		return
	}

	kb := engine.Window.Keyboard
	if kb.Down(keyboard.LeftCtrl) {
		// If left ctrl key is currently down, we apply scaling.
		sx, sy, sz := textNode.Scale()

		sx += math.Real(pos.X / 220)
		sz += math.Real(-pos.Y / 220)

		textNode.SetScale(sx, sy, sz)

	} else if kb.Down(keyboard.RightCtrl) {
		// If right ctrl key is currently down, we apply shearing.
		shx, shy, shz := textNode.Shear()

		shy += math.Real(pos.X / 220)
		shz += math.Real(-pos.Y / 220)

		textNode.SetShear(shx, shy, shz)

	} else if kb.Down(keyboard.LeftShift) || kb.Down(keyboard.RightShift) {
		// If an shift key is currently down, we apply relative rotation.
		pos.X /= 10
		pos.Y /= 10

		textNode.SetRelativeRot(textNode, math.Real(pos.Y), 0, math.Real(-pos.X))

	} else {
		// Otherwise we apply relative movement.
		x := math.Real(pos.X)
		z := math.Real(-pos.Y)
		textNode.SetRelativePos(textNode, x.Rounded(), 0, z.Rounded())
	}
}

// Event handler to reset cube transforms (their positions, scales, etc)
func resetTransforms(ev *event.Event) {
	log.Println("Reset")
	textNode.ResetTransform()
	textNode.SetPos(0, 0, -24)
}

// Event handler which toggles cursor grab
func toggleCursorGrabbed(ev *event.Event) {
	isGrabbed := engine.Window.CursorGrabbed()
	engine.Window.SetCursorGrabbed(!isGrabbed)
}

// Our scene graph will look like this:
//
// > Display
//     > Scene2d
//         > Camera2d
//         > textNode
//
func program() {
	var (
		err  error
		path = "src/code.google.com/p/azul3d/assets/fonts/vera/Vera.ttf"
	)
	text.DefaultOptions.Source, err = text.LoadFontFile(path)
	if err != nil {
		log.Fatal(err)
	}
	// Semi-transparent red background:
	//text.DefaultOptions.Background = color.RGBA{255, 0, 0, 64}
	text.DefaultOptions.Background = color.Transparent
	text.DefaultOptions.Size = 24

	text.Set(textNode, "Hello Azul3D!")
	textNode.SetParent(engine.Scene2d)
	textNode.SetPos(0, 0, -24)

	// Print scene graph
	engine.Renderer.PrintTree()

	// Grab the cursor
	engine.Window.SetCursorGrabbed(true)

	var stop func()
	stop = event.Define(event.Handlers{
		"keyboard-typed": func(ev *event.Event) {
			typedEvent := ev.Data.(*keyboard.TypedEvent)
			current := text.Get(textNode)
			text.Set(textNode, current+string(typedEvent.Rune))
		},

		// Listen for alt keys to toggle cursor grabbed
		"RightAlt": toggleCursorGrabbed,
		"LeftAlt":  toggleCursorGrabbed,

		// Listen for R key to reset transformations
		"R": func(ev *event.Event) {
			resetTransforms(ev)
			text.Set(textNode, "Hello Azul3D!")
		},

		// Listen for F key to try and free texture
		"F": func(ev *event.Event) {
			log.Println("text destroyed now (and GC ran)!")
			textNode.Destroy()
			runtime.GC()
		},

		// Listen for cursor position to move the sprite
		"cursor-position": onCursorPosition,

		"window-destroyed": func(ev *event.Event) {
			stop()
		},
	})
}

func main() {
	// Run our program, enter main loop.
	engine.Run(program)
}
