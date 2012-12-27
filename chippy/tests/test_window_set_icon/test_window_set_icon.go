// +build tests

package main

import(
    "code.google.com/p/azul3d/chippy"
    _ "image/png"
    "image"
    "time"
    "fmt"
    "os"
)

func main() {
    err := chippy.Init()
    if err != nil {
        panic(err.Error())
    }
    defer chippy.Destroy()

    defaultScreen, err := chippy.DefaultScreen()
    if err != nil {
        panic(err.Error())
    }

    configs, err := chippy.FrameBufferConfigs(defaultScreen)
    if err != nil {
        panic(err.Error())
    }

    config := chippy.ChooseConfig(configs, chippy.WorstConfig, chippy.BestConfig)
    if config == nil {
        panic("No frame buffer configuration found, Are your drivers up to date?")
    }

    fmt.Println("Setting window icon (2_128x128) (perfect)")
	// Open the file.
	file, err := os.Open("data/window_2_128x128.png")
	if err != nil {
        panic(err)
	}
	defer file.Close()

	// Decode the image.
	m, _, err := image.Decode(file)
	if err != nil {
        panic(err)
	}


    win, err := chippy.NewWindow(defaultScreen, config, m)
    if err != nil {
        panic(err.Error())
    }
    fmt.Println("Opened a window with these Frame Buffer configurations:")
    fmt.Println(win.FrameBufferConfig())
    contextVersion, err := win.ContextVersionString()
    if err != nil {
        panic(err.Error())
    }
    fmt.Println("The window is capable of OpenGL", contextVersion)

    time.Sleep(15 * time.Second)
}

