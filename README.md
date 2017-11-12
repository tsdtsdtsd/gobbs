# gobbs

[![Build Status](https://travis-ci.org/bwmarrin/discordgo.svg?branch=master)](https://travis-ci.org/bwmarrin/discordgo)
[![Go Report Card](https://goreportcard.com/badge/github.com/tsdtsdtsd/gobbs)](https://goreportcard.com/report/github.com/tsdtsdtsd/gobbs)
[![Godoc](https://godoc.org/github.com/tsdtsdtsd/gobbs?status.svg)](https://godoc.org/github.com/tsdtsdtsd/gobbs)

A [Blum-Blum-Shub-Generator](https://en.wikipedia.org/wiki/Blum_Blum_Shub) in Go.

## Status

Testing

## Usage

Create a new generator with default config:

```
import "github.com/tsdtsdtsd/gobbs"

g, err := gobbs.NewGenerator()
```

`NewGenerator()` will generate two random blum primes, a blum integer and a random seed. If you want to use your own numbers from an other source, you can use `NewGeneratorWithConfig()`:

```
g, err := gobbs.NewGeneratorWithConfig(&gobbs.Config{
    PrimeP: myPrimeOne,
    PrimeQ: myPrimeTwo,
    Seed: mySeed,
    Bits: 1024,
})
```

The generator implements `io.Reader`:

```
buf := make([]byte, 1)

for {
    g.Read(buf)
    fmt.Printf("%02x.", buf)
}
```

## Benchmarks

```
bits = 1024
readLength = 512

> go test -bench=. -benchmem -count 3
goos: windows
goarch: amd64
pkg: github.com/tsdtsdtsd/gobbs
BenchmarkInit-8               10         298040350 ns/op         2540071 B/op       8105 allocs/op
BenchmarkInit-8                3         399833433 ns/op         3571288 B/op      11109 allocs/op
BenchmarkInit-8                5         297086520 ns/op         2594611 B/op       8185 allocs/op
BenchmarkRead-8               50          37528676 ns/op         9440499 B/op      16393 allocs/op
BenchmarkRead-8               50          37291990 ns/op         9440641 B/op      16394 allocs/op
BenchmarkRead-8               50          38036620 ns/op         9440500 B/op      16393 allocs/op
PASS
ok      github.com/tsdtsdtsd/gobbs      13.857s
```

## Credits

Heavily inspired by (basically a port of) https://github.com/foolean/blum-blum-shub