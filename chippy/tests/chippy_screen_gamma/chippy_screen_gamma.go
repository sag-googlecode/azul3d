package main

import (
	"code.google.com/p/azul3d/chippy"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	log.SetFlags(0)

	// Enable debug messages
	chippy.SetDebugOutput(os.Stdout)

	err := chippy.Init()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer chippy.Destroy()

	screens := chippy.Screens()
	log.Printf("There are %d screens.\n", len(screens))
	log.Println("Default screen:", chippy.DefaultScreen())

	for i, screen := range screens {
		log.Printf("\nScreen %d - %s", i, screen)
	}

	log.Printf("\nEnter screen: #")
	var screen int
	_, err = fmt.Scanln(&screen)
	if err != nil {
		log.Fatal(err)
	}

	if screen < 0 || screen > len(screens)-1 {
		log.Fatal("Incorrect screen number.")
	}

	do := func(brightness, contrast, gamma float32) {
		fmt.Printf("brightness = %f, contrast = %f, gamma = %f\n", brightness, contrast, gamma)
		// do actual operation
		err := screens[screen].SetBrightnessContrastGamma(brightness, contrast, gamma)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("\nNormal")
	do(0.0, 1.0, 1.0)

	// Stuck on the wrong settings? uncomment this line and you'll get reset to normal here
	// return

	log.Println("\nBrightness")
	for brightness := float32(-1); brightness < 1.0; brightness += (1.0 / 30.0) {
		do(brightness, 1.0, 1.0)
		<-time.After(60 * time.Millisecond)
	}
	// do(-1.0, 1.0, 1.0)
	// do(-0.5, 1.0, 1.0)
	// do(0.5, 1.0, 1.0)
	// do(1.0, 1.0, 1.0)

	log.Println("\nContrast")
	for contrast := float32(0); contrast < 0.5; contrast += (0.5 / 30.0) {
		do(0.0, contrast, 1.0)
		<-time.After(60 * time.Millisecond)
	}
	// do(0.0, 0.0, 1.0)
	// do(0.0, 0.5, 1.0)
	// do(0.0, 10000.0, 1.0)

	log.Println("\nGamma")
	for gamma := float32(0.5); gamma < 5.0; gamma += (5.0 / 30.0) {
		do(0.0, 1.0, gamma)
		<-time.After(60 * time.Millisecond)
	}
	// do(0.0, 1.0, 5.0)
	// do(0.0, 1.0, 0.5)

	// Uncomment this line and the gamma will stay active on the screen after the program exits
	//screens[screen].SetAutoRestoreOriginalGammaRamp(false)
}
