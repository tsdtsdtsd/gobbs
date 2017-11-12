package main

import (
	"fmt"

	"github.com/tsdtsdtsd/gobbs"
)

func main() {

	g, _ := gobbs.NewGenerator()

	buf := make([]byte, 1)

	for {
		g.Read(buf)
		fmt.Printf("%02x.", buf)
	}

}
