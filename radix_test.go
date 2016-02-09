package radix

import (
	"sort"
	"testing"
)

func TestSorting(t *testing.T) {
	var (
		sizes  = []int{0, 1, 2, 3, 10, 100, 1000, 10000}
		sorter = map[string]func([]int32){
			"int32 radix sort MSB": Int32MSB,
			"int32 radix sort LSB": Int32LSB,
			"int32 standard sort":  stdSort,
		}
	)
	for _, size := range sizes {
		xs := pop(size)
		for desc, s := range sorter {
			ys := make([]int32, size)
			copy(ys, xs)
			s(ys)
			if !sort.IsSorted(byInt32(ys)) {
				t.Errorf("array of size %d was not correctly sorted by %s", size, desc)
			}
		}
	}
}

func Benchmark_Int32_RadixMSB_10(b *testing.B)     { benchmarkInt32(b, Int32MSB, 10) }
func Benchmark_Int32_RadixMSB_100(b *testing.B)    { benchmarkInt32(b, Int32MSB, 100) }
func Benchmark_Int32_RadixMSB_1000(b *testing.B)   { benchmarkInt32(b, Int32MSB, 1000) }
func Benchmark_Int32_RadixMSB_10000(b *testing.B)  { benchmarkInt32(b, Int32MSB, 10000) }
func Benchmark_Int32_RadixMSB_100000(b *testing.B) { benchmarkInt32(b, Int32MSB, 100000) }

func Benchmark_Int32_RadixLSB_10(b *testing.B)     { benchmarkInt32(b, Int32LSB, 10) }
func Benchmark_Int32_RadixLSB_100(b *testing.B)    { benchmarkInt32(b, Int32LSB, 100) }
func Benchmark_Int32_RadixLSB_1000(b *testing.B)   { benchmarkInt32(b, Int32LSB, 1000) }
func Benchmark_Int32_RadixLSB_10000(b *testing.B)  { benchmarkInt32(b, Int32LSB, 10000) }
func Benchmark_Int32_RadixLSB_100000(b *testing.B) { benchmarkInt32(b, Int32LSB, 100000) }

func Benchmark_Int32_StdSort_10(b *testing.B)     { benchmarkInt32(b, stdSort, 10) }
func Benchmark_Int32_StdSort_100(b *testing.B)    { benchmarkInt32(b, stdSort, 100) }
func Benchmark_Int32_StdSort_1000(b *testing.B)   { benchmarkInt32(b, stdSort, 1000) }
func Benchmark_Int32_StdSort_10000(b *testing.B)  { benchmarkInt32(b, stdSort, 10000) }
func Benchmark_Int32_StdSort_100000(b *testing.B) { benchmarkInt32(b, stdSort, 100000) }

func benchmarkInt32(b *testing.B, sorter func([]int32), size int) {
	ys := make([][]int32, b.N)
	for n := range ys {
		ys[n] = pop(size)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		sorter(ys[n])
	}
}

func stdSort(xs []int32) {
	sort.Sort(byInt32(xs))
}

var g = prng()

func pop(size int) []int32 {
	xs := make([]int32, size)
	for i := range xs {
		xs[i] = int32(g.next())
	}
	return xs
}

type byInt32 []int32

func (xs byInt32) Len() int           { return len(xs) }
func (xs byInt32) Swap(i, j int)      { xs[i], xs[j] = xs[j], xs[i] }
func (xs byInt32) Less(i, j int) bool { return xs[i] < xs[j] }
