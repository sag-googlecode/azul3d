// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package ui

import (
	"code.google.com/p/azul3d/math"
	"code.google.com/p/azul3d/scene"
	"code.google.com/p/azul3d/scene/color"
	"code.google.com/p/azul3d/scene/sprite"
	"sync"
)

type element struct {
	sync.RWMutex
	node, sprite *scene.Node
	states       StateOptions
	state        StateType
}

var elementTag = scene.NewProp("ui.element")

func getElement(n *scene.Node) *element {
	e, ok := n.Tag(elementTag)
	if !ok {
		return nil
	}
	return e.(*element)
}

func mustGetElement(n *scene.Node) *element {
	e, ok := n.Tag(elementTag)
	if !ok {
		panic("Specified node is not an sprite.")
	}
	return e.(*element)
}

func New(name string) *scene.Node {
	n := scene.New(name)
	e := new(element)

	n.SetTag(elementTag, e)
	e.node = n
	e.sprite = sprite.New(name + "-sprite")
	e.sprite.SetParent(e.node)

	// Center size
	//sprite.SetSize(e.sprite, 128, 128)
	//e.sprite.SetPos(64, 0, -64)

	return n
}

func SetOption(n *scene.Node, s StateType, o OptionType, value interface{}) {
	e := mustGetElement(n)
	e.Lock()
	defer e.Unlock()

	options, ok := e.states[s]
	if !ok {
		e.states[s] = Options{o: value}
	} else {
		options[o] = value
	}
}

func Option(n *scene.Node, s StateType, o OptionType) (value interface{}, ok bool) {
	e := mustGetElement(n)
	e.Lock()
	defer e.Unlock()

	options, ok := e.states[s]
	if !ok {
		return nil, false
	}
	value, ok = options[o]
	return
}

func Apply(n *scene.Node, s StateOptions) {
	err := s.Validate()
	if err != nil {
		panic(err)
	}

	e := mustGetElement(n)
	e.Lock()
	defer e.Unlock()

	if e.states == nil {
		e.states = s
	} else {
		for state, options := range s {
			_, ok := e.states[state]
			if !ok {
				e.states[state] = options
			} else {
				for opt, val := range options {
					e.states[state][opt] = val
				}
			}
		}
	}

	options, ok := s[e.state]
	var parent *scene.Node
	var width, height int
	c, colorScale := color.None, color.None
	if ok {
		for opt, val := range options {
			switch opt {
			case Parent:
				parent = val.(*scene.Node)
			case Width:
				width = val.(int)
			case Height:
				height = val.(int)
			case Color:
				c = val.(color.Color)
			case ColorScale:
				colorScale = val.(color.Color)
			}
		}
	}
	sprite.SetSize(e.sprite, math.Real(width), math.Real(height))
	if parent != nil {
		n.SetParent(parent)
	}
	if !c.Equals(color.None) {
		color.Set(e.sprite, c)
	}
	if !colorScale.Equals(color.None) {
		color.SetScale(e.sprite, colorScale)
	}
}

func SetState(n *scene.Node, s StateType) {
	e := mustGetElement(n)
	e.Lock()
	defer e.Unlock()

	e.state = s
}

func State(n *scene.Node) StateType {
	e := mustGetElement(n)
	e.RLock()
	defer e.RUnlock()

	return e.state
}

func Destroy(n *scene.Node) {
	e := mustGetElement(n)
	e.Lock()
	defer e.Unlock()

	e.node = nil
	e.sprite = nil
}
