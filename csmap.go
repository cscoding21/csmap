package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cscoding21/csmap/gen"
)

func main() {
	filePath := flag.String("f", "manifest.yaml", "path to the manifest file")
	flag.Parse()

	err := gen.Generate(*filePath)

	if err != nil {
		fmt.Println(fmt.Errorf("error generating code: %v", err))
		os.Exit(-1)
	}
}
