// +build examples

package main

import(
	"code.google.com/p/azul3d/scene/renderer"
	"code.google.com/p/azul3d/scene/texture"
	"code.google.com/p/azul3d/scene/sprite"
	"code.google.com/p/azul3d/scene/camera"
	"code.google.com/p/azul3d/scene/geom"
	"code.google.com/p/azul3d/scene"
	"code.google.com/p/azul3d/math"
	"code.google.com/p/azul3d/chippy"
	"code.google.com/p/azul3d/chippy/keyboard"
	"code.google.com/p/azul3d/event"
	"code.google.com/p/azul3d"
	_ "image/png"
	"runtime"
	"math/rand"
	"time"
	"log"
	"os"
)

var(
	// Create the engine.
	engine = azul3d.NewEngine()

	// root sprite
	root *scene.Node
)

// Event handler to manipulate cubes when cursor is moved
func onCursorPosition(ev *event.Event) {
	pos := ev.Data.(*chippy.CursorPositionEvent)

	// If the cursor is not grabbed, we do not transform cubes.
	if !engine.Window.CursorGrabbed() {
		return
	}

	kb := engine.Window.Keyboard
	if kb.Down(keyboard.RightCtrl) || kb.Down(keyboard.RightShift) {
		// If right ctrl key is currently down, we apply shearing to current
		// cube.
		shx, shy, shz := root.Shear()

		shy += math.Real(pos.X / 220)
		shz += math.Real(-pos.Y / 220)

		root.SetShear(shx, shy, shz)

	} else if kb.Down(keyboard.LeftShift) || kb.Down(keyboard.RightShift) {
		// If left ctrl key is currently down, we apply scaling to current
		// cube.
		sx, sy, sz := root.Scale()

		sx += math.Real(pos.X / 220)
		sz += math.Real(-pos.Y / 220)

		root.SetScale(sx, sy, sz)

	} else {
		// Otherwise we apply relative movement to the current cube.
		x := math.Real(pos.X)
		z := math.Real(-pos.Y)
		root.SetRelativePos(root, x.Rounded(), 0, z.Rounded())
	}
}

// Event handler to reset cube transforms (their positions, scales, etc)
func resetTransforms(ev *event.Event) {
	log.Println("Reset")
	root.ResetTransform()
}

func printViewed(ev *event.Event) {
	inView := camera.InView(engine.Camera2d, root.PosVec3(), root.Parent())

	log.Printf("Visibility: %s=%t\n", root.Name(), inView)
}

// Event handler which toggles cursor grab
func toggleCursorGrabbed(ev *event.Event) {
	isGrabbed := engine.Window.CursorGrabbed()
	engine.Window.SetCursorGrabbed(!isGrabbed)
}

func createSprite(parent *scene.Node, tex *texture.Texture2D) {
	s := sprite.New("sprite")

	// Center size
	sprite.SetSize(s, 16, 16)
	s.SetParent(parent)

	width, height := sprite.TotalSize(s)
	halfWidth := width / 2
	halfHeight := height / 2

	max, min := 640, 0
	x := math.Real(rand.Intn(max - min) + min)
	max, min = 480, 0
	y := math.Real(rand.Intn(max - min) + min)
	s.SetPos(x + halfWidth.Rounded(), s.PosVec3().Y, -y - halfHeight.Rounded())

	texture.Set(s, texture.DefaultLayer, tex)

	// Hide the sprite, we will show it once the texture is loaded.
	s.Hide()

	// Start an goroutine to print when the texture is loaded
	<-tex.LoadNotify()

	// Configure where the texture will show
	sprite.SetFrames(s, []texture.Region{
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

	sprite.SetFrameRate(s, (1000 / 24) * time.Millisecond)

	sprite.SetPlaying(s, true)

	s.Show()
}

func program() {
	root = engine.Scene2d.New("myrootNode")

	t, err := renderer.LoadTextureFile(engine.Renderer, "src/code.google.com/p/azul3d/assets/textures/player.png")
	if err != nil {
		log.Fatal(err)
	}

	// We know it's an 2D texture, which we want for the Texture2D.Region
	// method.
	tex := t.(*texture.Texture2D)

	// Grab the cursor
	engine.Window.SetCursorGrabbed(true)

	// We can benefit greatly by collecting all the like-sprites into a single
	// mesh. We can even do this at a shockingly fast rate, since sprites are
	// made of very few vertices.
	//
	// The flaw is that sprites are animated on the CPU by re-uploading their
	// vertex UV's after moving them. Instead of performing a Collect() call
	// every 24FPS, we could Collect() only once and animate the UV's on the
	// GPU using a shader.
	root.Hide()

	var newlyCollected = make(chan *scene.Node, 1)
	go func() {
		for{
			// Collect all the sprites into a single mesh.
			n, c := geom.Collect(root)
			log.Println("Collected", n, "nodes")

			// Wait until pre-frame to attach to graph.
			newlyCollected <- c
		}
	}()

	var collected *scene.Node
	event.Handle("pre-frame", func(e *event.Event) {
		var c *scene.Node
		select{
		case c = <-newlyCollected:
			break
		default:
			return
		}
		if collected != nil {
			// Destroy the previous collected mesh node.
			collected.Destroy()
		}
		collected = c

		// Attach the mesh to scene2d.
		collected.SetParent(engine.Scene2d)
	})

	var stop func()
	stop = event.Define(event.Handlers{
		// Listen for alt keys to toggle cursor grabbed
		"RightAlt": toggleCursorGrabbed,
		"LeftAlt": toggleCursorGrabbed,

		// Listen for R key to reset transformations
		"R": resetTransforms,

		// Listen for V key to print view status
		"V": printViewed,

		"keyboard-typed": func(ev *event.Event) {
			typedEvent := ev.Data.(*keyboard.TypedEvent)
			if typedEvent.Rune == rune(' ') {
				createSprite(root, tex)
			}
		},

		// Listen for F key to try and free texture
		"F": func(ev *event.Event) {
			log.Println("Sprite destroyed now (and GC ran)!")
			root.Destroy()
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
	rand.Seed(time.Now().Unix())

	// For debugging anything
	azul3d.SetDebugOutput(os.Stdout)

	// Initialize azul3d
	err := azul3d.Init()
	if err != nil {
		log.Fatal(err)
	}

	// Launch program
	go program()

	// Enter main loop
	azul3d.MainLoop()
}

