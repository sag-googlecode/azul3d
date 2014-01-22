package main

import (
	"azul3d.org/chippy/keyboard"
	"azul3d.org/engine"
	"azul3d.org/event"
	"azul3d.org/scene/color"
	"azul3d.org/scene/ui"
	_ "image/png"
	"log"
	"math/rand"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func program() {
	rand.Seed(time.Now().Unix())

	panel := ui.New("panel")
	panel.SetParent(engine.Scene2d)
	panel.SetPos(160, 0, -110)

	//panel.SetOption(ui.Width, 300)
	//panel.SetOption(ui.Height, 200)
	panel.SetOption(ui.Color, color.Hex("#FF0000"))

	var buttons []*ui.Element
	event.Handle("keyboard-typed", func(e *event.Event) {
		ev := e.Data.(*keyboard.TypedEvent)
		if ev.Rune == ' ' {
			button := ui.New("button")
			button.SetParent(panel.Node)

			button.SetOption(ui.Width, random(40, 100))
			button.SetOption(ui.Height, random(15, 26))
			button.SetOption(ui.Text, "Click me!")
			button.SetOption(ui.Overflow, ui.Hidden)
			button.SetOption(ui.Color, color.Color{
				R: rand.Float32(),
				G: rand.Float32(),
				B: rand.Float32(),
				A: 1.0,
			})
			buttons = append(buttons, button)

			panel.SetOption(ui.Color, color.Color{
				R: rand.Float32(),
				G: rand.Float32(),
				B: rand.Float32(),
				A: 1.0,
			})

		} else if ev.Rune == '1' {
			space := panel.Option(ui.LayoutSpace).(int)
			if space == 0 {
				log.Println("LayoutSpace=15px")
				panel.SetOption(ui.LayoutSpace, 15)
			} else if space == 15 {
				log.Println("LayoutSpace=26px")
				panel.SetOption(ui.LayoutSpace, 26)
			} else {
				log.Println("LayoutSpace=AutoSpace")
				panel.SetOption(ui.LayoutSpace, ui.AutoSpace)
			}

		} else if ev.Rune == '2' {
			layoutWrap := panel.Option(ui.LayoutWrap).(int)
			if layoutWrap == 0 {
				log.Println("LayoutWrap=250px")
				panel.SetOption(ui.LayoutWrap, 250)
			} else if layoutWrap == 250 {
				log.Println("LayoutWrap=400px")
				panel.SetOption(ui.LayoutWrap, 400)
			} else if layoutWrap == 400 {
				log.Println("LayoutWrap=NoWrap")
				panel.SetOption(ui.LayoutWrap, ui.NoWrap)
			} else {
				log.Println("LayoutWrap=AutoWrap")
				panel.SetOption(ui.LayoutWrap, ui.AutoWrap)
			}
		} else if ev.Rune == '3' {
			layout := panel.Option(ui.Layout).(ui.LayoutType)
			if layout == ui.Horizontal {
				log.Println("Layout=Vertical")
				panel.SetOption(ui.Layout, ui.Vertical)
			} else {
				log.Println("Layout=Horizontal")
				panel.SetOption(ui.Layout, ui.Horizontal)
			}
		}
	})
}

func main() {
	// Run our program, enter main loop.
	engine.Run(program)
}
