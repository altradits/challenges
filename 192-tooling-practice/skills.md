# Skills for 192-tooling-practice

## What You'll Learn

**Previous:** [191-p2pkh-script](../191-p2pkh-script/skills.md) | **Next:** none — you finished the course. Go contribute (see 158).

**Challenge:** Wield the four tools every Go contributor runs daily.

## 1. The Race Detector

```bash
go test -race ./174-worker-pool/ ./175-sync-atomic/
```

`-race` instruments memory accesses and reports concurrent read/write
pairs with both stack traces. It slows tests ~10x and needs cgo — that
is normal. LND and btcd CI run it; a data race in your PR fails review
automatically. If it reports, fix the sharing (mutex, atomic, channel),
never silence the test.

## 2. Fuzzing

A fuzz test states a *property* and lets Go hunt inputs that break it:

```go
func FuzzReverseTwice(f *testing.F) {
    f.Add([]byte("hello")) // seed corpus
    f.Fuzz(func(t *testing.T, b []byte) {
        in := append([]int(nil), bytesToInts(b)...)
        once := append([]int(nil), in...)
        ReverseInPlace(once)
        ReverseInPlace(once)
        if !reflect.DeepEqual(once, in) {
            t.Fatalf("not identity: %v", in)
        }
    })
}
```

```bash
go test -run FuzzReverseTwice ./184-two-pointers/   # seeds only
go test -fuzz FuzzReverseTwice -fuzztime 60s ./184-two-pointers/
```

Crashes land in `testdata/fuzz/` as regression tests — keep them. Write
your property against your own 184 solution.

## 3. Benchmarks and pprof

```bash
go test -run '^$' -bench . -benchtime 100x -cpuprofile cpu.out ./180-sorting-algorithms/
go tool pprof -top cpu.out
```

`-run '^$'` skips unit tests; `-bench .` runs all benchmarks. `-top`
shows where CPU time went — expect the sort inner loops, not the test
harness. Interactive mode (`go tool pprof cpu.out`, then `top`, `list
QuickSort`) digs deeper. Memory: `-memprofile mem.out`, then `-top`.
Rule: measure before optimizing, and keep the benchmark so the next
change cannot silently regress.

## 4. go vet

```bash
go vet ./...
```

`vet` flags real bugs: copied locks, `Printf` mismatches, unreachable
code, atomic copies. Zero findings is the bar — upstream CI enforces it.
Read each finding, understand the bug class, fix the code. `gofmt -l .`
alongside keeps formatting clean (empty output means clean).
