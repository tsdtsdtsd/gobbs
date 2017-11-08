# gobbs

Trying to implement a Blum-Blum-Shub-Generator in Go.

## Status

- MVP for generating 2 random primes, suitable for creation of a blum integer
- the blum integer
- currently benchmarking two different approaches

```
> go test -bench=. -count 3 -benchmem
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

## Credits

Heavily inspired by https://github.com/foolean/blum-blum-shub