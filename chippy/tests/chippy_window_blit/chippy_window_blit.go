// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.
// +build tests

// Test - Opens an windows, uses blitting to copy pixels on to it.
package main

import (
	"code.google.com/p/azul3d/chippy"
	"runtime/pprof"
	"image"
	"image/draw"
	_ "image/png"
	"log"
	"os"
	"time"
	"flag"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func program() {
	defer chippy.Exit()

	var err error

	// Load the image that we'll use for the window icon
	file, err := os.Open("src/code.google.com/p/azul3d/chippy/tests/data/chippy_720x320.png")
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	rgba, ok := img.(*image.RGBA)
	if !ok {
		// Need to convert to RGBA image
		b := img.Bounds()
		rgba = image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
		draw.Draw(rgba, rgba.Bounds(), img, b.Min, draw.Src)
	}

	window := chippy.NewWindow()
	window.SetSize(720, 320)
	window.SetTransparent(true)
	window.SetDecorated(false)
	window.SetAlwaysOnTop(true)

	// Actually open the windows
	screen := chippy.DefaultScreen()

	// Center the window on the screen
	window.SetPositionCenter(screen)

	err = window.Open(screen)
	if err != nil {
		log.Fatal(err)
	}

	// Print out what they currently has property-wise
	log.Println(window)

	events := window.Events()
	defer window.CloseEvents(events)

	var(
		measureBlitSpeed = time.After(10 * time.Second)
		measuringBlitSpeed bool
		numBlits int
		totalBlitTime time.Duration
	)

	for {
		// In order to clear a rectangle on the window, this is much faster
		// than PixelBlit() with an fully transparent image:
		//
		// window.PixelClear(image.Rect(0, 0, 30, 30))

		// Blit the image to the window, at x=0, y=0, blitting the entire image
		start := time.Now()
		window.PixelBlit(0, 0, rgba)
		blitTime := time.Since(start)
		log.Println("PixelBlit():", blitTime)

		numBlits++
		totalBlitTime += blitTime

		if measuringBlitSpeed {
			select{
			case e := <-events:
				switch e.(type) {
				case *chippy.CloseEvent:
					chippy.Exit()
					goto stats

				default:
					// We don't care about whatever event this is.
					break
				}
			default:
				break
			}
			continue
		}

		// Wait for an paint event
		gotPaintEvent := false
loop: for !gotPaintEvent {
			select {
			case <-measureBlitSpeed:
				measuringBlitSpeed = true
				break loop

			case e := <-events:
				switch e.(type) {
				case *chippy.PaintEvent:
					log.Println(e)
					gotPaintEvent = true

				case *chippy.CloseEvent:
					chippy.Exit()
					goto stats

				default:
					// We don't care about whatever event this is.
					break
				}
			}
		}
	}

stats:
	log.Printf("%d PixelBlit() over %v\n", numBlits, totalBlitTime)
	log.Printf("Average blit time: %v\n", totalBlitTime / time.Duration(numBlits))
}

func main() {
    flag.Parse()
    if *cpuprofile != "" {
        f, err := os.Create(*cpuprofile)
        if err != nil {
            log.Fatal(err)
        }
		defer f.Close()
        pprof.StartCPUProfile(f)
        defer pprof.StopCPUProfile()
    }

	log.SetFlags(0)

	// Enable debug output
	chippy.SetDebugOutput(os.Stdout)

	// Initialize Chippy
	err := chippy.Init()
	if err != nil {
		log.Fatal(err)
	}

	// Start program
	go program()

	// Enter main loop
	chippy.MainLoop()
}
