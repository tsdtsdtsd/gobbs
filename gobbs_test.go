package gobbs_test

import (
	"testing"

	"github.com/tsdtsdtsd/gobbs"
)

const bits = 7

func BenchmarkGetBlumA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, _, err := gobbs.GetBlumA(bits)
		if err != nil {
			panic(err)
		}
	}
}
func BenchmarkGetBlumB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, _, err := gobbs.GetBlumB(bits)
		if err != nil {
			panic(err)
		}
	}
}
