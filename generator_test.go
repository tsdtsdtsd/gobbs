package gobbs_test

import (
	"testing"

	"github.com/tsdtsdtsd/gobbs"
)

const bits = 64

func BenchmarkLoadStream(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g, _ := gobbs.NewGeneratorWithConfig(&gobbs.Config{
			Bits: bits,
		})
		_, _ = g.NewStream()
	}
}
