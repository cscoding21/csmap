package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cscoding21/csgen"
	"github.com/cscoding21/csmap/gen"
)

func main() {
	filePath := flag.String("f", "csmap.yaml", "path to the manifest file")
	flag.Parse()

	mcfg := csgen.GetDefaultPackageConfig()
	err := gen.Generate(*filePath, mcfg)

	if err != nil {
		fmt.Println(fmt.Errorf("error generating code: %v", err))
		os.Exit(-1)
	}
}
