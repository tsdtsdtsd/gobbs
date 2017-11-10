package main

import (
	"fmt"

	"github.com/tsdtsdtsd/gobbs"
)

func main() {

	g, _ := gobbs.NewGenerator()
	stream, _ := g.NewStream()

	for value := range stream.Start().C {
		fmt.Printf("%02x.", value)
	}

}
