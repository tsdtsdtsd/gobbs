# gobbs

[![Build Status](https://travis-ci.org/bwmarrin/discordgo.svg?branch=master)](https://travis-ci.org/bwmarrin/discordgo)
[![Godoc](https://godoc.org/github.com/tsdtsdtsd/gobbs?status.svg)](https://godoc.org/github.com/tsdtsdtsd/gobbs)
[![Go Report Card](https://goreportcard.com/badge/github.com/tsdtsdtsd/gobbs)](https://goreportcard.com/report/github.com/tsdtsdtsd/gobbs)

Trying to implement a [Blum-Blum-Shub-Generator](https://en.wikipedia.org/wiki/Blum_Blum_Shub) in Go.

## Status

Testing

## Tasks done and notes

- MVP for generating 2 random primes, suitable for creation of a blum integer
- Benchmarked two different algorythm approaches for the blum units. 
- Following benchmark was run with only a size of **7-bit** for the primes:

```
> go test -bench=. -benchmem -count 3 
goos: windows
goarch: amd64
pkg: github.com/tsdtsdtsd/gobbs
BenchmarkGetBlumA-8        10000            115685 ns/op           47646 B/op        449 allocs/op
BenchmarkGetBlumA-8        10000            119112 ns/op           48442 B/op        457 allocs/op
BenchmarkGetBlumA-8        10000            124850 ns/op           49142 B/op        463 allocs/op
BenchmarkGetBlumB-8        10000            137123 ns/op           56365 B/op        533 allocs/op
BenchmarkGetBlumB-8        10000            136516 ns/op           56226 B/op        532 allocs/op
BenchmarkGetBlumB-8        10000            139084 ns/op           57232 B/op        541 allocs/op
PASS
ok      github.com/tsdtsdtsd/gobbs      7.854s
```

- Slowly structuring the package. 
- Method for random seeds added.
- This benchmark was run with a size of **64-bit** for the primes:

```
> go test -bench=. -benchmem -count 5 
goos: windows
goarch: amd64
pkg: github.com/tsdtsdtsd/gobbs
BenchmarkGetBlumUnits-8             5000            260299 ns/op           39657 B/op        153 allocs/op
BenchmarkGetBlumUnits-8             5000            252194 ns/op           38859 B/op        152 allocs/op
BenchmarkGetBlumUnits-8             5000            264566 ns/op           39508 B/op        153 allocs/op
BenchmarkGetBlumUnits-8             5000            264938 ns/op           39656 B/op        153 allocs/op
BenchmarkGetBlumUnits-8             5000            257945 ns/op           39920 B/op        154 allocs/op
BenchmarkGetBlumA-8                 2000            760240 ns/op          117883 B/op        456 allocs/op
BenchmarkGetBlumA-8                 2000            787551 ns/op          119651 B/op        464 allocs/op
BenchmarkGetBlumA-8                 2000            787917 ns/op          119020 B/op        463 allocs/op
BenchmarkGetBlumA-8                 2000            766809 ns/op          118607 B/op        456 allocs/op
BenchmarkGetBlumA-8                 2000            841421 ns/op          118671 B/op        464 allocs/op
BenchmarkRandomSeeds-8            300000              4034 ns/op             584 B/op         15 allocs/op
BenchmarkRandomSeeds-8            300000              4032 ns/op             584 B/op         15 allocs/op
BenchmarkRandomSeeds-8            300000              4037 ns/op             584 B/op         15 allocs/op
BenchmarkRandomSeeds-8            300000              4040 ns/op             584 B/op         15 allocs/op
BenchmarkRandomSeeds-8            300000              4063 ns/op             584 B/op         15 allocs/op
PASS
ok      github.com/tsdtsdtsd/gobbs      21.326s
```

- Finished generator
- Added Stream concept

## Credits

Heavily inspired by (basically a port of) https://github.com/foolean/blum-blum-shub