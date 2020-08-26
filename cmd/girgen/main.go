package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/diamondburned/handy/internal/gir"
)

func main() {
	flag.Parse()

	var girPath = flag.Arg(0)
	if girPath == "" {
		log.Fatalln("Missing .gir path. Usage: girgen file.gir.")
	}

	if err := gir.ParseRepositoryFile(girPath); err != nil {
		log.Fatalln(err)
	}

	gen := gir.NewGotk3Generator("handy")
	gir.SetActiveNamespace(0).GenerateToFile(gen)

	// Process the filename.
	outputPath := strings.Split(girPath, ".")[0]
	outputPath = strings.ToLower(outputPath)
	// Optionally trim the version dash.
	outputPath = strings.Split(outputPath, "-")[0]

	f, err := os.Create(fmt.Sprintf("%s_generated.go", outputPath))
	if err != nil {
		log.Fatalln("Failed to create output file:", err)
	}
	defer f.Close()

	fmt.Fprintf(f, "%#v", gen)
}
