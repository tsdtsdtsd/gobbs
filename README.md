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

## Credits

Heavily inspired by (basically a port of) https://github.com/foolean/blum-blum-shub