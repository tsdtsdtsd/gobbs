package gobbs_test

import (
	"testing"

	"github.com/tsdtsdtsd/gobbs"
)

const bits = 1024
const readLength = 512

func BenchmarkRead(b *testing.B) {

	g, _ := gobbs.NewGeneratorWithConfig(&gobbs.Config{
		Bits: bits,
	})
	buf := make([]byte, bits)

	for i := 0; i < b.N; i++ {
		g.Read(buf)
	}
}
