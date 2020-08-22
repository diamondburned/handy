package main

import (
	"fmt"
	"log"

	"github.com/diamondburned/handy/gir"
)

func main() {
	r, err := gir.ParseRepositoryFile("./Handy-1.gir")
	if err != nil {
		log.Fatalln(err)
	}

	f := gir.NewGotk3Generator("handy")
	r.Namespaces[0].GenerateToFile(f)

	fmt.Printf("%#v\n", f)
}
