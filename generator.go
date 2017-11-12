package gobbs

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// Generator defines a Blum-Blum-Shub generator
type Generator struct {
	p    *big.Int
	q    *big.Int
	n    *big.Int
	seed *big.Int
	x0   *big.Int
	x1   *big.Int

	bits     int
	readInit bool
}

// NewGenerator returns a Generator with default config
func NewGenerator() (*Generator, error) {
	return NewGeneratorWithConfig(DefaultConfig)
}

// NewGeneratorWithConfig returns a Generator with given config
func NewGeneratorWithConfig(config *Config) (*Generator, error) {

	g := &Generator{
		p:    big.NewInt(0),
		q:    big.NewInt(0),
		n:    big.NewInt(0),
		seed: big.NewInt(0),
		x0:   big.NewInt(0),
		x1:   big.NewInt(0),
		bits: config.Bits,
	}

	// Handle config

	if config.PrimeP != nil && config.PrimeQ != nil {
		// @todo validate primes first
		g.p.Set(config.PrimeP)
		g.q.Set(config.PrimeQ)
		g.n.Mul(g.p, g.q)
	}

	if config.Seed != nil {
		g.seed.Set(config.Seed)
	}

	// Check values

	if g.n.Cmp(bigZero) == 0 {
		_, _, _, err := g.CalcBlumUnits()
		if err != nil {
			return g, err
		}
	}

	if g.seed.Cmp(bigZero) == 0 {
		_, err := g.CalcRandomSeed()
		if err != nil {
			return g, err
		}
	}

	return g, nil
}

func (g *Generator) Read(p []byte) (n int, err error) {
	if l := len(p); l > 0 {
		for n < l {
			p[n] = g.readByte()
			n++
		}
	}
	return
}

// CalcBlumUnits calculates, sets to *g and returns two Blum primes (p) and (q), as well as their product, the Blum integer (n).
func (g *Generator) CalcBlumUnits() (*big.Int, *big.Int, *big.Int, error) {

	var (
		err error
		tmp = big.NewInt(0)
	)

	g.q = big.NewInt(0)
	g.p, err = rand.Prime(rand.Reader, g.bits)
	if err != nil {
		return g.p, g.q, g.n, err
	}

	for tmp.Mod(g.n, bigFour).Cmp(bigOne) != 0 || g.p.Cmp(g.q) == 0 {
		g.q, err = rand.Prime(rand.Reader, g.bits)
		if err != nil {
			return g.p, g.q, g.n, err
		}

		g.n.Mul(g.p, g.q)
	}

	return g.p, g.q, g.n, nil
}

// CalcRandomSeed finds, sets to *g and returns a random large integer (x), that is relatively prime to (n).
func (g *Generator) CalcRandomSeed() (*big.Int, error) {

	if g.n.Cmp(bigZero) == 0 {
		return nil, fmt.Errorf("no valid blum integer (n) set")
	}

	var (
		err error

		r      = big.NewInt(0)
		maxInt = big.NewInt(int64(g.bits))
	)

	g.seed = big.NewInt(0)

	// maxInt = 2^n - 1 (n=bits)
	maxInt.Exp(bigTwo, maxInt, nil)
	maxInt.Sub(maxInt, bigOne)

	if maxInt.Cmp(bigZero) == 0 {
		return g.seed, fmt.Errorf("error calculating maximum integer size")
	}

	for r.Cmp(bigOne) != 0 {
		g.seed, err = rand.Int(rand.Reader, maxInt)
		if err != nil {
			return g.seed, err
		}
		r.GCD(nil, nil, g.n, g.seed)
	}

	return g.seed, err
}

func (g *Generator) readByte() byte {

	var (
		bitCounter uint
		val        uint
	)

	if !g.readInit {
		// x0 = (seed ^ 2) mod n
		g.x0.Exp(g.seed, bigTwo, g.n)
		g.readInit = true
	}

	for {
		// x1 = (x0 ^ 2) mod n
		g.x1.Exp(g.x0, bigTwo, g.n)

		if bitCounter == 0 {
			val = 0
		}

		if g.x1.Bit(0) == 1 {
			val |= (1 << bitCounter)
		}

		bitCounter++
		g.x0.Set(g.x1)

		if bitCounter == 8 {
			return byte(val)
		}
	}

}
