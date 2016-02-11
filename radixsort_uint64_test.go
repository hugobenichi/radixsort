package radixsort

import (
	"sort"
	"testing"
)

func TestUint64Sorting(t *testing.T) {
	var (
		sizes  = []int{0, 1, 2, 3, 10, 1e2, 1e3, 1e4, 1e5, 1e6}
		sorter = map[string]func([]uint64){
			"uint64 radix sort MSD": Uint64MSD,
			"uint64 radix sort LSD": Uint64LSD,
			"uint64 standard sort":  uint64_stdSort,
		}
	)
	for _, size := range sizes {
		xs := uint64_pop(size)
		for desc, s := range sorter {
			ys := make([]uint64, size)
			copy(ys, xs)
			s(ys)
			if !sort.IsSorted(byUint64(ys)) {
				t.Errorf("array of size %d was not correctly sorted by %s", size, desc)
			}
		}
	}
}

func Benchmark_Uint64_RadixMSD_100(b *testing.B)    { benchmarkUint64(b, Uint64MSD, 100) }
func Benchmark_Uint64_RadixMSD_1000(b *testing.B)   { benchmarkUint64(b, Uint64MSD, 1000) }
func Benchmark_Uint64_RadixMSD_10000(b *testing.B)  { benchmarkUint64(b, Uint64MSD, 10000) }
func Benchmark_Uint64_RadixMSD_100000(b *testing.B) { benchmarkUint64(b, Uint64MSD, 100000) }

func Benchmark_Uint64_RadixLSD_100(b *testing.B)    { benchmarkUint64(b, Uint64LSD, 100) }
func Benchmark_Uint64_RadixLSD_1000(b *testing.B)   { benchmarkUint64(b, Uint64LSD, 1000) }
func Benchmark_Uint64_RadixLSD_10000(b *testing.B)  { benchmarkUint64(b, Uint64LSD, 10000) }
func Benchmark_Uint64_RadixLSD_100000(b *testing.B) { benchmarkUint64(b, Uint64LSD, 100000) }

func Benchmark_Uint64_StandardSort_100(b *testing.B)    { benchmarkUint64(b, uint64_stdSort, 100) }
func Benchmark_Uint64_StandardSort_1000(b *testing.B)   { benchmarkUint64(b, uint64_stdSort, 1000) }
func Benchmark_Uint64_StandardSort_10000(b *testing.B)  { benchmarkUint64(b, uint64_stdSort, 10000) }
func Benchmark_Uint64_StandardSort_100000(b *testing.B) { benchmarkUint64(b, uint64_stdSort, 100000) }

func Benchmark_Uint64_Insertion_100(b *testing.B) { benchmarkUint64(b, uint64_insertion, 100) }

func benchmarkUint64(b *testing.B, sorter func([]uint64), size int) {
	ys := make([][]uint64, b.N)
	for n := range ys {
		ys[n] = uint64_pop(size)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		sorter(ys[n])
	}
}

func uint64_stdSort(xs []uint64) {
	sort.Sort(byUint64(xs))
}

func uint64_pop(size int) []uint64 {
	xs := make([]uint64, size)
	for i := range xs {
		xs[i] = uint64(g.next())
	}
	return xs
}

type byUint64 []uint64

func (xs byUint64) Len() int           { return len(xs) }
func (xs byUint64) Swap(i, j int)      { xs[i], xs[j] = xs[j], xs[i] }
func (xs byUint64) Less(i, j int) bool { return xs[i] < xs[j] }
