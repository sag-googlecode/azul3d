package main

import "code.google.com/p/azul3d"
import "fmt"

func destroy() {
	fmt.Println("destroy() called")
}

func init() {
	fmt.Println("init() called")
}

func main() {
	defer azul.Destroy()
	// The above 'defer azul.Destroy()' line is required
	// so that all registered destroy functions will be
	// called upon the shutdown of the program.

	azul.RegisterDestroy(destroy)
}
