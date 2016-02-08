package radix

import (
	"math/rand"
	"sort"
	"testing"
)

func TestSorting(t *testing.T) {
	data := [][]int32{
		pop([]int32{}),
		pop(make([]int32, 10)),
		pop(make([]int32, 1000)),
		pop(make([]int32, 100000)),
	}

	sorter := []func([]int32){
		Int32MSB,
		Int32LSB,
		stdSort,
		//insertion,  // too slow
		//shell,			// too slow
	}

	for _, d := range data {
		for _, s := range sorter {
			xs := make([]int32, len(d))
			copy(xs, d)
			s(xs)
			if !sort.IsSorted(byInt32(xs)) {
				t.Errorf("was not sorted")
			}
		}
	}
}

var (
	xs10     = pop(make([]int32, 10))
	xs100    = pop(make([]int32, 100))
	xs1000   = pop(make([]int32, 1000))
	xs10000  = pop(make([]int32, 10000))
	xs100000 = pop(make([]int32, 1000000))
)

func Benchmark_Int32_RadixMSB_10(b *testing.B)     { benchmarkInt32(b, Int32MSB, xs10) }
func Benchmark_Int32_RadixMSB_100(b *testing.B)    { benchmarkInt32(b, Int32MSB, xs100) }
func Benchmark_Int32_RadixMSB_1000(b *testing.B)   { benchmarkInt32(b, Int32MSB, xs1000) }
func Benchmark_Int32_RadixMSB_10000(b *testing.B)  { benchmarkInt32(b, Int32MSB, xs10000) }
func Benchmark_Int32_RadixMSB_100000(b *testing.B) { benchmarkInt32(b, Int32MSB, xs100000) }

func Benchmark_Int32_RadixLSB_10(b *testing.B)     { benchmarkInt32(b, Int32LSB, xs10) }
func Benchmark_Int32_RadixLSB_100(b *testing.B)    { benchmarkInt32(b, Int32LSB, xs100) }
func Benchmark_Int32_RadixLSB_1000(b *testing.B)   { benchmarkInt32(b, Int32LSB, xs1000) }
func Benchmark_Int32_RadixLSB_10000(b *testing.B)  { benchmarkInt32(b, Int32LSB, xs10000) }
func Benchmark_Int32_RadixLSB_100000(b *testing.B) { benchmarkInt32(b, Int32LSB, xs100000) }

func Benchmark_Int32_StdSort_10(b *testing.B)     { benchmarkInt32(b, stdSort, xs10) }
func Benchmark_Int32_StdSort_100(b *testing.B)    { benchmarkInt32(b, stdSort, xs100) }
func Benchmark_Int32_StdSort_1000(b *testing.B)   { benchmarkInt32(b, stdSort, xs1000) }
func Benchmark_Int32_StdSort_10000(b *testing.B)  { benchmarkInt32(b, stdSort, xs10000) }
func Benchmark_Int32_StdSort_100000(b *testing.B) { benchmarkInt32(b, stdSort, xs100000) }

func benchmarkInt32(b *testing.B, sorter func([]int32), xs []int32) {
	ys := make([]int32, len(xs))
	copy(ys, xs)
	for n := 0; n < b.N; n++ {
		sorter(xs)
	}
}

func stdSort(xs []int32) {
	sort.Sort(byInt32(xs))
}

func pop(xs []int32) []int32 {
	for i := range xs {
		xs[i] = int32(rand.Int())
		if rand.Int()&1 == 1 {
			xs[i] = ^xs[i]
		}
	}
	return xs
}

type byInt32 []int32

func (xs byInt32) Len() int           { return len(xs) }
func (xs byInt32) Swap(i, j int)      { xs[i], xs[j] = xs[j], xs[i] }
func (xs byInt32) Less(i, j int) bool { return xs[i] < xs[j] }
