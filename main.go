package main

import (
	"fmt"
	"os"

	"github.com/gausszhou/softfetch/internal/detect"
	"github.com/gausszhou/softfetch/internal/display"
)

var (
	version   = "dev"
	commit    = "unknown"
	buildDate = "unknown"
)

func main() {
	detectors := detect.GetCoreDetectors()
	result := detect.Detect(detectors...)

	display.PrintResult(result)

	if len(os.Args) > 1 {
		if os.Args[1] == "--version" || os.Args[1] == "-v" {
			fmt.Printf("SoftFetch version: %s\n", version)
		}
	}
}
