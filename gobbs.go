// Package gobbs provides a Blum-Blum-Shub generator.
// See: https://en.wikipedia.org/wiki/Blum_Blum_Shub
package gobbs

import (
	"math/big"
)

var (
	bigZero  = big.NewInt(0) // These should be consts, but consts can't be structs :(
	bigOne   = big.NewInt(1)
	bigTwo   = big.NewInt(2)
	bigThree = big.NewInt(3)
	bigFour  = big.NewInt(4)
)
