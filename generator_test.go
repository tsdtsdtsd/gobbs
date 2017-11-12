package gobbs_test

import (
	"testing"

	"github.com/tsdtsdtsd/gobbs"
)

const bits = 1024
const readLength = 512

var g *gobbs.Generator

func BenchmarkInit(b *testing.B) {

	for i := 0; i < b.N; i++ {
		g, _ = gobbs.NewWithConfig(&gobbs.Config{
			Bits: bits,
		})
	}
}

func BenchmarkRead(b *testing.B) {

	// g, _ := gobbs.NewWithConfig(&gobbs.Config{
	// 	Bits: bits,
	// })
	buf := make([]byte, bits)

	for i := 0; i < b.N; i++ {
		g.Read(buf)
	}
}
