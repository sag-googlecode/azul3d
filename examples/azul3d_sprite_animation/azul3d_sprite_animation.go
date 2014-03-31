package main

import (
	"azul3d.org/v0/chippy"
	"azul3d.org/v0/chippy/keyboard"
	"azul3d.org/v0/engine"
	"azul3d.org/v0/event"
	"azul3d.org/v0/math"
	"azul3d.org/v0/scene"
	"azul3d.org/v0/scene/bucket"
	"azul3d.org/v0/scene/camera"
	"azul3d.org/v0/scene/renderer"
	"azul3d.org/v0/scene/sprite"
	"azul3d.org/v0/scene/texture"
	_ "image/png"
	"log"
	"runtime"
	"time"
)

var (
	// Player sprite
	player *scene.Node
)

// Event handler to manipulate cubes when cursor is moved
func onCursorPosition(ev *event.Event) {
	pos := ev.Data.(chippy.CursorPositionEvent)

	// If the cursor is not grabbed, we do not transform cubes.
	if !engine.Window.CursorGrabbed() {
		return
	}

	kb := engine.Window.Keyboard
	if kb.Down(keyboard.LeftCtrl) {
		// If left ctrl key is currently down, we apply scaling to current
		// cube.
		sx, sy, sz := player.Scale()

		sx += float64(pos.X / 220)
		sz += float64(-pos.Y / 220)

		player.SetScale(sx, sy, sz)

	} else if kb.Down(keyboard.RightCtrl) {
		// If right ctrl key is currently down, we apply shearing to current
		// cube.
		shx, shy, shz := player.Shear()

		shy += float64(pos.X / 220)
		shz += float64(-pos.Y / 220)

		player.SetShear(shx, shy, shz)

	} else if kb.Down(keyboard.LeftShift) || kb.Down(keyboard.RightShift) {
		// If an shift key is currently down, we apply relative rotation to the
		// current cube.
		width, height := sprite.Size(player)
		width += float32(pos.X)
		height -= float32(pos.Y)
		if width <= 0 {
			width = 1
		}
		if height <= 0 {
			height = 1
		}
		width = float32(math.Rounded(float64(width)))
		height = float32(math.Rounded(float64(height)))
		sprite.SetSize(player, width, height)

	} else {
		// Otherwise we apply relative movement to the current cube.
		x := float64(pos.X)
		z := float64(-pos.Y)
		player.SetRelativePos(player, math.Rounded(x), 0, math.Rounded(z))
	}
}

// Event handler to reset cube transforms (their positions, scales, etc)
func resetTransforms(ev *event.Event) {
	log.Println("Reset")
	player.ResetTransform()

	width, height := sprite.TotalSize(player)
	halfWidth := width / 2
	halfHeight := height / 2
	player.SetPos(math.Rounded(float64(halfWidth)), player.PosVec3().Y, -math.Rounded(float64(halfHeight)))
}

func printViewed(ev *event.Event) {
	inView := camera.InView(engine.Camera2d, player.PosVec3(), player.Parent())

	log.Printf("Visibility: %s=%t\n", player.Name(), inView)
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
//         > Player
//
func program() {
	// Create player sprite
	player = sprite.New("Player")

	// Center size
	sprite.SetSize(player, 256, 256) // Actual size of the sprite is 16x16, but we will scale it up here
	player.SetParent(engine.Scene2d)

	width, height := sprite.TotalSize(player)
	halfWidth := width / 2
	halfHeight := height / 2
	player.SetPos(math.Rounded(float64(halfWidth)), player.PosVec3().Y, -math.Rounded(float64(halfHeight)))
	bucket.Set(player, bucket.Transparent)

	t, err := renderer.LoadTextureFile(engine.Renderer, "src/azul3d.org/v0/assets/textures/player.png")
	if err != nil {
		log.Fatal(err)
	}
	t.SetMinFilter(texture.Nearest)
	t.SetMagFilter(texture.Nearest)

	// We know it's an 2D texture, which we want for the Texture2D.Region
	// method.
	tex := t.(*texture.Texture2D)
	texture.Set(player, texture.DefaultLayer, tex)

	// Hide the player, we will show it once the texture is loaded.
	player.Hide()

	// Start an goroutine to print when the texture is loaded
	go func() {
		log.Println("Loading the texture...")
		<-tex.LoadNotify()
		log.Println("Loaded the texture.")

		// Configure where the texture will show
		sprite.SetFrames(player, []texture.Region{
			tex.Region(0, 0, 16, 16),
			tex.Region(16, 0, 32, 16),
			tex.Region(32, 0, 48, 16),
			tex.Region(48, 0, 64, 16),

			tex.Region(0, 16, 16, 32),
			tex.Region(16, 16, 32, 32),
			tex.Region(32, 16, 48, 32),
			tex.Region(48, 16, 64, 32),

			tex.Region(0, 32, 16, 48),
			tex.Region(16, 32, 32, 48),
			tex.Region(32, 32, 48, 48),
			tex.Region(48, 32, 64, 48),

			tex.Region(0, 48, 16, 64),
			tex.Region(16, 48, 32, 64),
			tex.Region(32, 48, 48, 64),
			tex.Region(48, 48, 64, 64),
		})

		sprite.SetFrameRate(player, (1000/24)*time.Millisecond)

		sprite.SetPlaying(player, true)

		player.Show()
	}()

	// Print scene graph
	engine.Renderer.PrintTree()

	// Grab the cursor
	engine.Window.SetCursorGrabbed(true)

	var stop func()
	stop = event.Define(event.Handlers{
		// Listen for alt keys to toggle cursor grabbed
		"RightAlt": toggleCursorGrabbed,
		"LeftAlt":  toggleCursorGrabbed,

		// Listen for R key to reset transformations
		"R": resetTransforms,

		// Listen for V key to print view status
		"V": printViewed,

		// Listen for F key to try and free texture
		"F": func(ev *event.Event) {
			log.Println("Sprite destroyed now (and GC ran)!")
			player.Detatch()
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
