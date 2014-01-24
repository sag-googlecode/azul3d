package main

import (
	"azul3d.org/engine"
	"azul3d.org/event"
	"azul3d.org/math"
	"azul3d.org/scene"
	"azul3d.org/scene/bucket"
	"azul3d.org/scene/geom"
	"azul3d.org/scene/geom/procedural"
	"azul3d.org/scene/renderer"
	"azul3d.org/scene/texture"
	"azul3d.org/scene/transparency"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"sync"
	"time"
)

var (
	// Cards.
	center, one, two, three *scene.Node

	// Lock for when we are changing any of the below variables
	globalLock sync.RWMutex
)

func createCard(name string) *scene.Node {
	// Create an geom node
	geomNode := scene.New(name)

	// Build the card geom
	card := procedural.Card(-3, 3, -3, 3, texture.Region{0, 0, 1, 1}, geom.Static)

	// Attach it to the geom.Node
	geom.Add(geomNode, card)

	// Enable alpha blending transparency on the node
	transparency.Set(geomNode, transparency.AlphaBlend)

	// Load texture file
	//
	// You might try one of the following textures instead to increase load
	// time and show asynchronous file loading better.
	//
	// Note that the larger the texture the more memory it will consume, and it
	// could eat all your system's memory.
	prefix := "src/azul3d.org/assets/textures"
	//texFilePath := prefix + "/clouds_8192x8192.jpg"
	//texFilePath := prefix+"/clouds_4096x4096.jpg"
	//texFilePath := prefix+"/clouds_2048x2048.jpg"
	//texFilePath := prefix+"/clouds_1024x1024.jpg"
	texFilePath := prefix + "/texture_coords_1024x1024.png"

	tex, err := renderer.LoadTextureFile(engine.Renderer, texFilePath)
	if err != nil {
		log.Fatal(err)
	}

	// Assign texture to the card
	texture.Set(geomNode, texture.DefaultLayer, tex)

	return geomNode
}

// Event handler to reset card transforms (their positions, scales, etc)
func resetTransforms(ev *event.Event) {
	log.Println("Reset")

	center.ResetTransform()
	center.SetPos(0, 10, 0)
}

// Event handler which toggles cursor grab
func toggleCursorGrabbed(ev *event.Event) {
	isGrabbed := engine.Window.CursorGrabbed()
	engine.Window.SetCursorGrabbed(!isGrabbed)
}

// Event handler which toggles transparency
func toggleBucket(ev *event.Event) {
	switch bucket.Get(one) {
	case bucket.Background:
		bucket.Set(one, bucket.Opaque)
	case bucket.Opaque:
		bucket.Set(one, bucket.Transparent)
	case bucket.Transparent:
		bucket.Set(one, bucket.Sorted)
	case bucket.Sorted:
		bucket.Set(one, bucket.Unsorted)
	case bucket.Unsorted:
		bucket.Set(one, bucket.Background)
	}
	log.Println(bucket.Get(one))
}

// Our scene graph will look like this:
//
// > Display
//     > Scene2d
//         > Camera2d
//         > One
//             > Two
//                 > Three
//
func program() {
	center = engine.Scene3d.New("center")
	center.SetPos(0, 10, 0)

	one = createCard("one")
	one.SetParent(center)
	one.SetPos(0, -4, 0)
	bucket.Set(one, bucket.Transparent)

	two = createCard("two")
	two.SetParent(one)
	two.SetPos(0, 4, 0)

	three = createCard("three")
	three.SetParent(two)
	three.SetPos(0, 4, 0)

	// Only used when bucket == Sorted
	bucket.SetSort(one, 100)
	bucket.SetSort(two, 333)
	bucket.SetSort(three, 57)

	// Otherwise they would each inherit one's sort value because they're
	// children nodes of one.
	bucket.SetSortForced(one, true)
	bucket.SetSortForced(two, true)
	bucket.SetSortForced(three, true)

	// Print scene graph
	engine.Renderer.PrintTree()

	var stop func()
	stop = event.Define(event.Handlers{
		// Listen for T key to change buckets
		"T": toggleBucket,

		"pre-frame": func(ev *event.Event) {
			// Rotate at 45 deg/sec
			deltaSeconds := float32(renderer.Clock(engine.Renderer).Delta()) / float32(time.Second)
			center.SetRelativeRot(center, 0, 0, math.Real(45*deltaSeconds))
		},

		"window-destroyed": func(ev *event.Event) {
			stop()
		},
	})
}

func main() {
	// Run our program, enter main loop.
	engine.Run(program)
}
