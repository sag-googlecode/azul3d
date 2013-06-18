// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.
//
// Code generated by this program is also under
// the above license.

// Generation program to assist with OpenGL wrapper creation
package main

import (
	"code.google.com/p/azul3d/native/opengl/generator/generator"
	"flag"
	"log"
)

var (
	specFilePath string
	outputDir    string
)

func init() {
	log.SetFlags(0)
	flag.StringVar(&specFilePath, "spec", "", "JSON specification file")
	flag.StringVar(&outputDir, "o", "", "output directory")
}

func main() {
	flag.Parse()

	if len(specFilePath) == 0 || len(outputDir) == 0 {
		flag.Usage()
	}

	generator.Generate(specFilePath, outputDir)
}
