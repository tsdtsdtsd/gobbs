package gobbs

import "math/big"

// DefaultConfig contains some good starter settings
var DefaultConfig = &Config{
	Bits: 1024,
}

// Config contains configuration settings for a generator
type Config struct {
	PrimeP *big.Int
	PrimeQ *big.Int
	Seed   *big.Int

	Bits int
}
