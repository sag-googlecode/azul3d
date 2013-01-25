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

		for i, mode := range screen.ScreenModes() {
			log.Printf("    -> ScreenMode %d - %s", i, mode)
		}
	}

	fmt.Printf("Change Screen: #")
	var screen int
	_, err = fmt.Scanln(&screen)
	if err != nil {
		log.Fatal(err)
	}

	if screen < 0 || screen > len(screens)-1 {
		log.Fatal("Incorrect screen number.")
	}

	fmt.Printf("Change Screen #%d to mode: #", screen)
	var mode int
	_, err = fmt.Scanln(&mode)
	if err != nil {
		log.Fatal(err)
	}

	if mode < 0 || mode > len(screens[screen].ScreenModes())-1 {
		log.Fatal("Incorrect screen number.")
	}

	// Change screen mode
	screens[screen].ScreenModes()[mode].Use()

	log.Println("Waiting 15 seconds...")
	<-time.After(15 * time.Second)

	// Disable this with screen.SetAutoRestoreOriginalScreenMode(false)
	log.Println("Original resolution should restore now")
}
