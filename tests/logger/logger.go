package main

import "code.google.com/p/azul3d/logger"
import "code.google.com/p/azul3d"

func main() {
	defer azul.Destroy()
	log := logger.New("MyLogger")
	log.Log("My message!", "Now a number:", 1)
}
