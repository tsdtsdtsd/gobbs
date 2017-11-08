package gobbs

import (
	"crypto/rand"
	"math/big"
)

func GetBlumA(bits int) (*big.Int, *big.Int, *big.Int, error) {

	var (
		p, q *big.Int
		err  error
		one  = big.NewInt(1)
		four = big.NewInt(4)
		n    = big.NewInt(0)
		tmp  = big.NewInt(0)
	)

	p, err = rand.Prime(rand.Reader, bits)
	if err != nil {
		return p, q, n, err
	}

	for tmp.Mod(n, four).Cmp(one) != 0 || p.Cmp(q) == 0 {
		q, err = rand.Prime(rand.Reader, bits)
		if err != nil {
			return p, q, n, err
		}

		n.Mul(p, q)
	}

	return p, q, n, nil
}

func GetBlumB(bits int) (*big.Int, *big.Int, *big.Int, error) {

	var (
		err error

		p    = big.NewInt(0)
		q    = big.NewInt(0)
		n    = big.NewInt(0)
		m    = big.NewInt(3 % 4)
		four = big.NewInt(4)
		tmp  = big.NewInt(0)
	)

	for tmp.Mod(p, four).Cmp(m) != 0 {
		p, err = rand.Prime(rand.Reader, bits)
		if err != nil {
			return p, q, n, err
		}
	}

	for tmp.Mod(q, four).Cmp(m) != 0 || p.Cmp(q) == 0 {
		q, err = rand.Prime(rand.Reader, bits)
		if err != nil {
			return p, q, n, err
		}
	}

	n.Mul(p, q)

	return p, q, n, nil
}
