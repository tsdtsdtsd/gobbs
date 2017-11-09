// Package gobbs provides a Blum-Blum-Shub generator.
package gobbs

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
)

var (
	bigZero = big.NewInt(0) // These should be consts, but consts can't be structs :(
	bigOne  = big.NewInt(1)
	bigTwo  = big.NewInt(2)
	bigFour = big.NewInt(4)

	// DefaultConfig contains some good starter settings
	DefaultConfig = &Config{
		Bits: 512,
	}
)

// Generator defines a Blum-Blum-Shub generator
type Generator struct {
	sync.Mutex
	p *big.Int
	q *big.Int
	n *big.Int

	bits int
}

// Config contains configuration settings for a generator
type Config struct {
	Bits int
}

// NewGenerator returns a ready to run Generator with default config
func NewGenerator() *Generator {
	return NewGeneratorWithConfig(DefaultConfig)
}

// NewGeneratorWithConfig returns a ready to run Generator with given config
func NewGeneratorWithConfig(config *Config) *Generator {
	g := &Generator{
		p:    big.NewInt(0),
		q:    big.NewInt(0),
		n:    big.NewInt(0),
		bits: config.Bits,
	}

	return g
}

// Run is temporary
func (g *Generator) Run() {
	p, q, n, err := g.GetBlumUnits()
	if err != nil {
		panic(err)
	}

	fmt.Println("p:", p)
	fmt.Println("q:", q)
	fmt.Println("n:", n)

	s, err := g.GetRandomSeed()
	if err != nil {
		panic(err)
	}

	fmt.Println("seed:", s)
}

// GetBlumUnits calculates and returns two Blum primes (p) and (q), as well as their product, the Blum integer (n).
func (g *Generator) GetBlumUnits() (*big.Int, *big.Int, *big.Int, error) {

	var (
		err error
		tmp = big.NewInt(0)
	)

	g.Lock()
	defer g.Unlock()

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

// GetRandomSeed finds and returns a random large integer (x), that is relatively prime to (n).
func (g *Generator) GetRandomSeed() (*big.Int, error) {

	var seed = big.NewInt(0)

	if g.n.Cmp(bigZero) == 0 {
		return seed, fmt.Errorf("no valid blum integer (n) set")
	}

	var (
		err error

		r      = big.NewInt(0)
		maxInt = big.NewInt(int64(g.bits))
	)

	// maxInt = 2^n - 1 (n=bits)
	maxInt.Exp(bigTwo, maxInt, nil)
	maxInt.Sub(maxInt, bigOne)

	if maxInt.Cmp(bigZero) == 0 {
		return seed, fmt.Errorf("error calculating maximum integer size")
	}

	for r.Cmp(bigOne) != 0 {
		seed, err = rand.Int(rand.Reader, maxInt)
		if err != nil {
			return seed, err
		}
		r.GCD(nil, nil, g.n, seed)
	}

	return seed, err
}
