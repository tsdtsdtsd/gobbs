package main

import (
	"fmt"

	"github.com/tsdtsdtsd/gobbs"
)

func main() {

	g, _ := gobbs.New()

	buf := make([]byte, 1)

	for {
		g.Read(buf)
		fmt.Printf("%02x.", buf)
	}

	// for {
	// 	g.Read(buf)
	// 	if int(buf[0]) >= 0 && int(buf[0]) <= 3 {
	// 		fmt.Printf("%d\n", int(buf[0]))
	// 		break
	// 	}
	// }

}
