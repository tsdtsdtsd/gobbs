package gobbs_test

import (
	"math/big"
	"testing"

	"github.com/tsdtsdtsd/gobbs"
)

const bits = 1024
const readLength = 512

var (
	g  *gobbs.Generator
	gn *big.Int

	bigOne   = big.NewInt(1)
	bigThree = big.NewInt(3)
	bigFour  = big.NewInt(4)
)

func TestNew(t *testing.T) {

	var err error
	g, err = gobbs.New()
	if err != nil {
		t.Error("Error creating generator: ", err)
	}
}

func TestCalcBlumUnits(t *testing.T) {

	var tmp = big.Int{}

	p, q, n, err := g.CalcBlumUnits()
	if err != nil {
		t.Error("Error calculating blum units: ", err)
	}

	// check n
	if tmp.Mul(p, q).Cmp(n) != 0 {
		t.Error("n is not a product of p*q")
	}

	// check if p & q are primes
	if !p.ProbablyPrime(1) {
		t.Error("p is not a prime")
	}
	if !q.ProbablyPrime(1) {
		t.Error("q is not a prime")
	}

	// check if p & q are congruent to 3 (mod 4)
	modCheck := big.NewInt(0)
	modCheck.Mod(bigThree, bigFour)

	if tmp.Mod(p, bigFour).Cmp(modCheck) != 0 {
		t.Error("p is not a blum prime")
	}
	if tmp.Mod(q, bigFour).Cmp(modCheck) != 0 {
		t.Error("q is not a blum prime")
	}

	gn = n

}

func TestCalcRandomSeed(t *testing.T) {

	var tmp = big.Int{}

	seed, err := g.CalcRandomSeed()
	if err != nil {
		t.Error("Error calculating random seed: ", err)
	}

	if tmp.GCD(nil, nil, gn, seed).Cmp(bigOne) != 0 {
		t.Error("GCD(n,seed) != 1 ")
	}
}

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
	buf := make([]byte, readLength)

	for i := 0; i < b.N; i++ {
		g.Read(buf)
	}
}
