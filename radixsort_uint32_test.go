package radixsort

import (
	"sort"
	"testing"
)

func TestUint32Sorting(t *testing.T) {
	var (
		sizes  = []int{0, 1, 2, 3, 10, 1e2, 1e3, 1e4, 1e5, 1e6}
		sorter = map[string]func([]uint32){
			"uint32 radix sort MSD": Uint32MSD,
			"uint32 radix sort LSD": Uint32LSD,
			"uint32 standard sort":  uint32_stdSort,
		}
	)
	for _, size := range sizes {
		xs := uint32_pop(size)
		for desc, s := range sorter {
			ys := make([]uint32, size)
			copy(ys, xs)
			s(ys)
			if !sort.IsSorted(byUint32(ys)) {
				t.Errorf("array of size %d was not correctly sorted by %s", size, desc)
			}
		}
	}
}

func Benchmark_Uint32_RadixMSD_100(b *testing.B)    { benchmarkUint32(b, Uint32MSD, 100) }
func Benchmark_Uint32_RadixMSD_1000(b *testing.B)   { benchmarkUint32(b, Uint32MSD, 1000) }
func Benchmark_Uint32_RadixMSD_10000(b *testing.B)  { benchmarkUint32(b, Uint32MSD, 10000) }
func Benchmark_Uint32_RadixMSD_100000(b *testing.B) { benchmarkUint32(b, Uint32MSD, 100000) }

func Benchmark_Uint32_RadixLSD_100(b *testing.B)    { benchmarkUint32(b, Uint32LSD, 100) }
func Benchmark_Uint32_RadixLSD_1000(b *testing.B)   { benchmarkUint32(b, Uint32LSD, 1000) }
func Benchmark_Uint32_RadixLSD_10000(b *testing.B)  { benchmarkUint32(b, Uint32LSD, 10000) }
func Benchmark_Uint32_RadixLSD_100000(b *testing.B) { benchmarkUint32(b, Uint32LSD, 100000) }

func Benchmark_Uint32_StandardSort_100(b *testing.B)    { benchmarkUint32(b, uint32_stdSort, 100) }
func Benchmark_Uint32_StandardSort_1000(b *testing.B)   { benchmarkUint32(b, uint32_stdSort, 1000) }
func Benchmark_Uint32_StandardSort_10000(b *testing.B)  { benchmarkUint32(b, uint32_stdSort, 10000) }
func Benchmark_Uint32_StandardSort_100000(b *testing.B) { benchmarkUint32(b, uint32_stdSort, 100000) }

func Benchmark_Uint32_Insertion_100(b *testing.B) { benchmarkUint32(b, uint32_insertion, 100) }

func benchmarkUint32(b *testing.B, sorter func([]uint32), size int) {
	ys := make([][]uint32, b.N)
	for n := range ys {
		ys[n] = uint32_pop(size)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		sorter(ys[n])
	}
}

func benchmarkUint32OneDigit(b *testing.B, sorter func([]uint32), size int) {
	ys := make([][]uint32, b.N)
	for n := range ys {
		ys[n] = uint32_pop(size)
		for i := range ys[n] {
			ys[n][i] = ys[n][i] & 0xFF
		}
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		sorter(ys[n])
	}
}

func uint32_stdSort(xs []uint32) {
	sort.Sort(byUint32(xs))
}

func uint32_pop(size int) []uint32 {
	xs := make([]uint32, size)
	for i := range xs {
		xs[i] = uint32(g.next())
	}
	return xs
}

type byUint32 []uint32

func (xs byUint32) Len() int           { return len(xs) }
func (xs byUint32) Swap(i, j int)      { xs[i], xs[j] = xs[j], xs[i] }
func (xs byUint32) Less(i, j int) bool { return xs[i] < xs[j] }
