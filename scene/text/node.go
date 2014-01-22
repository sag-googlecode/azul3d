// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package text

import (
	"azul3d.org/event"
	"azul3d.org/math"
	"azul3d.org/scene"
	"azul3d.org/scene/geom"
	"azul3d.org/scene/texture"
	"azul3d.org/scene/transparency"
	"image"
	"sync"
)

var (
	// Global atlas variables
	gaAccess    sync.RWMutex
	gaGlyphs    map[rune]*GlyphOptions
	gaTexture   *texture.Texture2D
	gaRasterMap map[rune]map[*GlyphOptions]*GlyphRaster
	gaRebuild   chan chan bool

	PTextObject = scene.NewProp("TextObject")
)

func globalAtlasBuilder() {
	for {
		var chs []chan bool

		complete := <-gaRebuild
		chs = append(chs, complete)

		length := len(gaRebuild)
		for i := 0; i < length; i++ {
			complete := <-gaRebuild
			chs = append(chs, complete)
		}

		gaAccess.Lock()

		if len(gaGlyphs) > 0 {
			var img *image.RGBA
			img, gaRasterMap = Atlas(gaGlyphs, 1, true)

			if img != nil {
				gaTexture = texture.New()
				gaTexture.SetImage(img)
				gaTexture.SetWrapModeU(texture.Clamp)
				gaTexture.SetWrapModeV(texture.Clamp)
				gaTexture.SetCompressed(false)
			}
		}

		gaAccess.Unlock()

		for _, ch := range chs {
			ch <- true
		}
	}
}

func init() {
	gaRebuild = make(chan chan bool)
	gaGlyphs = make(map[rune]*GlyphOptions)

	go globalAtlasBuilder()
}

type textObject struct {
	sync.RWMutex
	text           string
	node, textNode *scene.Node
	updated        bool
}

func (t *textObject) updateLater() {
	t.updated = false

	var stop func()
	stop = event.Handle("pre-frame", func(ev *event.Event) {
		t.Lock()
		defer t.Unlock()

		if t.node != nil {
			if !t.updated {
				t.update()
				t.updated = true
			}
		}

		stop()
	})
}

func (t *textObject) update() {
	gaAccess.Lock()
	for _, r := range t.text {
		gaGlyphs[r] = DefaultOptions
	}
	gaAccess.Unlock()

	complete := make(chan bool, 1)
	gaRebuild <- complete
	<-complete

	gaAccess.RLock()
	defer gaAccess.RUnlock()

	var (
		verts []geom.Vertex
		tcs   []texture.Coord
	)
	origin := math.Real(0)

	var last rune
	for index, r := range t.text {
		raster := gaRasterMap[r][DefaultOptions]
		if raster == nil {
			continue
		}

		sz := raster.Image.Bounds().Size()
		bearingX := math.Real(raster.HMetrics.BearingX)
		bearingY := math.Real(raster.HMetrics.BearingY)

		left := float32(origin + bearingX)
		right := float32(origin + bearingX + math.Real(sz.X))
		bottom := float32(-(math.Real(sz.Y) - bearingY))
		top := float32(bearingY)
		//log.Println(bottom, top)

		if index > 0 {
			x, y := Kerning(last, r, DefaultOptions)
			left += float32(x)
			bottom += float32(y)
		}

		origin += math.Real(raster.HMetrics.Advance)

		a := raster.Area
		rg := gaTexture.Region(a.Min.X, a.Min.Y, a.Max.X, a.Max.Y)

		// Bottom-Left triangle
		verts = append(verts, geom.Vertex{left, 0, top})
		verts = append(verts, geom.Vertex{left, 0, bottom})
		verts = append(verts, geom.Vertex{right, 0, bottom})
		tcs = append(tcs, texture.UV(rg.U, rg.V))
		tcs = append(tcs, texture.UV(rg.U, rg.T))
		tcs = append(tcs, texture.UV(rg.S, rg.T))

		// Top-Right triangle
		verts = append(verts, geom.Vertex{left, 0, top})
		verts = append(verts, geom.Vertex{right, 0, bottom})
		verts = append(verts, geom.Vertex{right, 0, top})
		tcs = append(tcs, texture.UV(rg.U, rg.V))
		tcs = append(tcs, texture.UV(rg.S, rg.T))
		tcs = append(tcs, texture.UV(rg.S, rg.V))

		last = r
	}

	// Create new node
	updatedText := scene.New("Text")
	geom.Add(updatedText, &geom.Mesh{
		Hint:     geom.Dynamic,
		Vertices: verts,
		TextureCoords: [][]texture.Coord{
			tcs,
		},
	})

	if t.textNode != nil {
		t.textNode.Destroy()
	}
	t.textNode = updatedText
	t.textNode.SetParent(t.node)

	if gaTexture == nil {
		texture.Remove(t.textNode, texture.DefaultLayer)
	} else {
		texture.Set(t.textNode, texture.DefaultLayer, gaTexture)
	}
}

func (t *textObject) set(n *scene.Node, text string) {
	t.Lock()
	defer t.Unlock()

	if t.text == text {
		return
	}
	t.text = text
	t.updateLater()
}

func (t *textObject) get() string {
	t.RLock()
	defer t.RUnlock()

	return t.text
}

func newTextObject(n *scene.Node) *textObject {
	t := new(textObject)
	t.node = n
	t.node.SetProp(PTextObject, t)
	return t
}

func getObj(n *scene.Node) *textObject {
	i, ok := n.Prop(PTextObject)
	if !ok {
		return nil
	}
	return i.(*textObject)
}

// [foo]This sounds [u]great [/foo][b]because...[/u] it is. [[this has brackets]].[/b]
func Set(n *scene.Node, text string) {
	transparency.Set(n, transparency.Multisample)
	i, ok := n.Prop(PTextObject)
	if !ok {
		o := newTextObject(n)
		i = o
	}
	o := i.(*textObject)
	o.set(n, text)
}

func Get(n *scene.Node) string {
	o := getObj(n)
	if o == nil {
		return ""
	}
	return o.get()
}

func Destroy(n *scene.Node) {
	o := getObj(n)
	if o == nil {
		return
	}
	o.node.Destroy()
}
