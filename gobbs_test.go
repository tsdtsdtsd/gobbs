package gobbs_test

import (
	"fmt"
	"testing"

	"github.com/tsdtsdtsd/gobbs"
)

const bits = 64

func BenchmarkGetBlumUnits(b *testing.B) {
	g := gobbs.NewGeneratorWithConfig(&gobbs.Config{
		Bits: bits,
	})

	for i := 0; i < b.N; i++ {
		_, _, _, err := g.GetBlumUnits()
		if err != nil {
			fmt.Println("ERROR:", err)
		}
	}
}

func BenchmarkRandomSeeds(b *testing.B) {
	g := gobbs.NewGeneratorWithConfig(&gobbs.Config{
		Bits: bits,
	})
	_, _, _, err := g.GetBlumUnits()
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	for i := 0; i < b.N; i++ {
		_, err := g.GetRandomSeed()
		if err != nil {
			fmt.Println("ERROR:", err)
		}
	}
}
