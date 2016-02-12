package radixsort

import (
	"sort"
	"testing"
)

func TestIntSorting(t *testing.T) {
	var (
		sizes  = []int{0, 1, 2, 3, 10, 1e2, 1e3, 1e4, 1e5, 1e6}
		sorter = map[string]func([]int){
			"int radix":         Int,
			"int standard sort": sort.Ints,
		}
	)
	for _, size := range sizes {
		xs := int_pop(size)
		for desc, s := range sorter {
			ys := make([]int, size)
			copy(ys, xs)
			s(ys)
			if !sort.IntsAreSorted(ys) {
				t.Errorf("array of size %d was not correctly sorted by %s", size, desc)
			}
		}
	}
}

func Benchmark_Int_RadixMSD_100(b *testing.B)    { benchmarkInt(b, Int, 100) }
func Benchmark_Int_RadixMSD_1000(b *testing.B)   { benchmarkInt(b, Int, 1000) }
func Benchmark_Int_RadixMSD_10000(b *testing.B)  { benchmarkInt(b, Int, 10000) }
func Benchmark_Int_RadixMSD_100000(b *testing.B) { benchmarkInt(b, Int, 100000) }

func Benchmark_Int_StandardSort_100(b *testing.B)    { benchmarkInt(b, sort.Ints, 100) }
func Benchmark_Int_StandardSort_1000(b *testing.B)   { benchmarkInt(b, sort.Ints, 1000) }
func Benchmark_Int_StandardSort_10000(b *testing.B)  { benchmarkInt(b, sort.Ints, 10000) }
func Benchmark_Int_StandardSort_100000(b *testing.B) { benchmarkInt(b, sort.Ints, 100000) }

func benchmarkInt(b *testing.B, sorter func([]int), size int) {
	ys := make([][]int, b.N)
	for n := range ys {
		ys[n] = int_pop(size)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		sorter(ys[n])
	}
}

func int_pop(size int) []int {
	xs := make([]int, size)
	for i := range xs {
		xs[i] = int(g.next())
	}
	return xs
}
