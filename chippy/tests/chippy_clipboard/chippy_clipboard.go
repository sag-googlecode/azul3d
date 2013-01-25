package main

import (
	"code.google.com/p/azul3d/chippy"
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

	finished := false
	for !finished {
		if chippy.ClipboardHasData() {
			finished = true

			contents, err := chippy.ClipboardString()
			if err != nil {
				log.Println("Clipboard string:", contents)
				chippy.SetClipboardString("Hello Chippy World!")
			} else {
				log.Fatal(err)
			}
		}

		if !finished {
			log.Println("Waiting 3 seconds for clipboard contents...")
			<-time.After(3 * time.Second)
		}
	}

	// Disable this with screen.SetAutoRestoreOriginalScreenMode(false)
	log.Println("Clearing clipboard contents after 15 seconds...")
	<-time.After(15 * time.Second)

	err = chippy.ClearClipboard()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Clipboar cleared")
}
