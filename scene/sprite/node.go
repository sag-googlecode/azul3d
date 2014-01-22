// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package sprite

import (
	"azul3d.org/event"
	"azul3d.org/math"
	"azul3d.org/scene"
	"azul3d.org/scene/geom"
	"azul3d.org/scene/geom/procedural"
	"azul3d.org/scene/texture"
	"azul3d.org/scene/transparency"
	"sync"
	"time"
)

type obj struct {
	sync.RWMutex
	node    *scene.Node
	updated bool
	mesh    *geom.Mesh
	patches *procedural.NinePatches
	regions *procedural.NineRegions

	frames           []texture.Region
	currentFrame     int
	frameRate        time.Duration
	stopPlaying      chan bool
	playing          bool
	cancelLastUpdate chan bool
}

func (o *obj) update() {
	geom.Remove(o.node, o.mesh)
	o.mesh = procedural.Ninepatch(o.patches, o.regions, geom.Dynamic)

	o.mesh.MakePixelPerfect()

	geom.Add(o.node, o.mesh)
}

func (o *obj) updateLater() {
	o.updated = false

	var stop func()
	stop = event.Handle("pre-frame", func(ev *event.Event) {
		o.Lock()
		defer o.Unlock()

		if o.node != nil {
			if !o.updated {
				o.update()
				o.updated = true
			}
		}
		stop()
	})
}

var PSpriteObject = scene.NewProp("SpriteObject")

func getObj(n *scene.Node) *obj {
	o, ok := n.Prop(PSpriteObject)
	if !ok {
		return nil
	}
	return o.(*obj)
}

func mustGetObj(n *scene.Node) *obj {
	o := getObj(n)
	if o != nil {
		return o
	}
	panic("Specified node is not an sprite.")
}

func New(name string) *scene.Node {
	n := scene.New(name)
	transparency.Set(n, transparency.Multisample)
	o := new(obj)
	o.stopPlaying = make(chan bool, 1)

	o.patches = &procedural.NinePatches{
		Width:  100,
		Height: 100,
	}

	o.regions = &procedural.NineRegions{
		Center: texture.Region{0, 0, 1, 1},
	}

	n.SetProp(PSpriteObject, o)
	o.node = n
	return n
}

func Destroy(n *scene.Node) {
	SetPlaying(n, false)

	o := mustGetObj(n)
	o.Lock()
	defer o.Unlock()

	o.node = nil
	o.mesh = nil
}

func SetSize(n *scene.Node, width, height math.Real) {
	if width <= 0 || height <= 0 {
		panic("SetSize(): Width and height must be greater than zero!")
	}

	o := mustGetObj(n)
	o.Lock()
	defer o.Unlock()

	o.updateLater()
	o.patches.Width = float32(width)
	o.patches.Height = float32(height)
}

func Size(n *scene.Node) (width, height math.Real) {
	o := mustGetObj(n)
	o.RLock()
	defer o.RUnlock()

	return math.Real(o.patches.Width), math.Real(o.patches.Height)
}

func TotalSize(n *scene.Node) (width, height math.Real) {
	o := mustGetObj(n)
	o.RLock()
	defer o.RUnlock()

	w := o.patches.Width + o.patches.Left + o.patches.Right
	h := o.patches.Height + o.patches.Bottom + o.patches.Top
	return math.Real(w), math.Real(h)
}

func SetBorders(n *scene.Node, left, right, bottom, top math.Real) {
	o := mustGetObj(n)
	o.Lock()
	defer o.Unlock()

	o.updateLater()
	o.patches.Left = float32(left)
	o.patches.Right = float32(right)
	o.patches.Bottom = float32(bottom)
	o.patches.Top = float32(top)
}

func Borders(n *scene.Node) (left, right, bottom, top math.Real) {
	o := mustGetObj(n)
	o.RLock()
	defer o.RUnlock()

	p := o.patches
	return math.Real(p.Left), math.Real(p.Right), math.Real(p.Bottom), math.Real(p.Top)
}

func ClearBorders(n *scene.Node) {
	SetBorders(n, 0, 0, 0, 0)
}

func SetTextureRegions(n *scene.Node, regions *procedural.NineRegions) {
	o := mustGetObj(n)
	o.Lock()
	defer o.Unlock()

	o.updateLater()

	cpy := *regions
	o.regions = &cpy
}

func TextureRegions(n *scene.Node) *procedural.NineRegions {
	o := mustGetObj(n)
	o.RLock()
	defer o.RUnlock()

	cpy := *o.regions
	return &cpy
}

func SetFrames(n *scene.Node, frames []texture.Region) {
	o := mustGetObj(n)
	o.Lock()
	defer o.Unlock()

	o.frames = make([]texture.Region, len(frames))
	copy(o.frames, frames)
}

func Frames(n *scene.Node) []texture.Region {
	o := mustGetObj(n)
	o.Lock()
	defer o.Unlock()

	cpy := make([]texture.Region, len(o.frames))
	copy(cpy, o.frames)
	return cpy
}

func SetFrameRate(n *scene.Node, rate time.Duration) {
	o := mustGetObj(n)
	o.Lock()
	defer o.Unlock()

	o.frameRate = rate
}

func FrameRate(n *scene.Node) time.Duration {
	o := mustGetObj(n)
	o.RLock()
	defer o.RUnlock()

	return o.frameRate
}

func play(n *scene.Node) {
	o := mustGetObj(n)

	for {
		select {
		case <-time.After(FrameRate(n)):
			// Swap the frame
			SetFrame(n, Frame(n)+1)

		case <-o.stopPlaying:
			return
		}
	}
}

func SetPlaying(n *scene.Node, playing bool) {
	o := mustGetObj(n)
	o.Lock()
	defer o.Unlock()

	if o.playing != playing {
		o.playing = playing
		if !playing {
			o.stopPlaying <- true
		} else {
			go play(n)
		}
	}
}

func Playing(n *scene.Node) bool {
	o := mustGetObj(n)
	o.RLock()
	defer o.RUnlock()

	return o.playing
}

func SetFrame(n *scene.Node, frame int) {
	o := mustGetObj(n)
	o.Lock()

	for frame >= len(o.frames) {
		frame -= len(o.frames)
	}
	if frame < 0 {
		frame = 0
	}
	o.currentFrame = frame

	o.Unlock()

	SetTextureRegions(n, &procedural.NineRegions{
		Center: FrameRegion(n, Frame(n)),
	})
}

func Frame(n *scene.Node) int {
	o := mustGetObj(n)
	o.RLock()
	defer o.RUnlock()

	return o.currentFrame
}

func FrameRegion(n *scene.Node, frame int) texture.Region {
	o := mustGetObj(n)
	o.RLock()
	defer o.RUnlock()

	for frame >= len(o.frames) {
		frame -= len(o.frames)
	}
	if frame < 0 {
		frame = 0
	}
	return o.frames[frame]
}
