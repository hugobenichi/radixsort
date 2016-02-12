package radixsort

import (
	"sort"
	"testing"
)

func TestUintSorting(t *testing.T) {
	var (
		sizes  = []int{0, 1, 2, 3, 10, 1e2, 1e3, 1e4, 1e5, 1e6}
		sorter = map[string]func([]uint){
			"uint radix":         Uint,
			"uint standard sort": uint_stdSort,
		}
	)
	for _, size := range sizes {
		xs := uint_pop(size)
		for desc, s := range sorter {
			ys := make([]uint, size)
			copy(ys, xs)
			s(ys)
			if !sort.IsSorted(byUint(ys)) {
				t.Errorf("array of size %d was not correctly sorted by %s", size, desc)
			}
		}
	}
}

func Benchmark_Uint_Radix_100(b *testing.B)    { benchmarkUint(b, Uint, 100) }
func Benchmark_Uint_Radix_1000(b *testing.B)   { benchmarkUint(b, Uint, 1000) }
func Benchmark_Uint_Radix_10000(b *testing.B)  { benchmarkUint(b, Uint, 10000) }
func Benchmark_Uint_Radix_100000(b *testing.B) { benchmarkUint(b, Uint, 100000) }

func Benchmark_Uint_StandardSort_100(b *testing.B)    { benchmarkUint(b, uint_stdSort, 100) }
func Benchmark_Uint_StandardSort_1000(b *testing.B)   { benchmarkUint(b, uint_stdSort, 1000) }
func Benchmark_Uint_StandardSort_10000(b *testing.B)  { benchmarkUint(b, uint_stdSort, 10000) }
func Benchmark_Uint_StandardSort_100000(b *testing.B) { benchmarkUint(b, uint_stdSort, 100000) }

func benchmarkUint(b *testing.B, sorter func([]uint), size int) {
	ys := make([][]uint, b.N)
	for n := range ys {
		ys[n] = uint_pop(size)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		sorter(ys[n])
	}
}

func uint_stdSort(xs []uint) {
	sort.Sort(byUint(xs))
}

func uint_pop(size int) []uint {
	xs := make([]uint, size)
	for i := range xs {
		xs[i] = uint(g.next())
	}
	return xs
}

type byUint []uint

func (xs byUint) Len() int           { return len(xs) }
func (xs byUint) Swap(i, j int)      { xs[i], xs[j] = xs[j], xs[i] }
func (xs byUint) Less(i, j int) bool { return xs[i] < xs[j] }
