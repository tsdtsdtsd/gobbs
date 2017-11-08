package main

import "github.com/tsdtsdtsd/gobbs"
import "fmt"

func main() {

	p, q, n, err := gobbs.GetBlum(7)
	if err != nil {
		panic(err)
	}

	fmt.Println("p1", p)
	fmt.Println("q1", q)
	fmt.Println("n1", n)

	fmt.Println("===============================")

	p, q, n, err = gobbs.GetPrimes(7)
	if err != nil {
		panic(err)
	}

	fmt.Println("p2", p)
	fmt.Println("q2", q)
	fmt.Println("n2", n)
}
