package radix

import (
	"sort"
	"testing"
)

func TestInt64Sorting(t *testing.T) {
	var (
		sizes  = []int{0, 1, 2, 3, 10, 1e2, 1e3, 1e4, 1e5, 1e6}
		sorter = map[string]func([]int64){
			"int64 radix sort MSD": Int64MSD,
			"int64 radix sort LSD": Int64LSD,
			"int64 standard sort":  int64_stdSort,
		}
	)
	for _, size := range sizes {
		xs := int64_pop(size)
		for desc, s := range sorter {
			ys := make([]int64, size)
			copy(ys, xs)
			s(ys)
			if !sort.IsSorted(byInt64(ys)) {
				t.Errorf("array of size %d was not correctly sorted by %s", size, desc)
			}
		}
	}
}

func Benchmark_Int64_RadixMSD_100(b *testing.B)    { benchmarkInt64(b, Int64MSD, 100) }
func Benchmark_Int64_RadixMSD_1000(b *testing.B)   { benchmarkInt64(b, Int64MSD, 1000) }
func Benchmark_Int64_RadixMSD_10000(b *testing.B)  { benchmarkInt64(b, Int64MSD, 10000) }
func Benchmark_Int64_RadixMSD_100000(b *testing.B) { benchmarkInt64(b, Int64MSD, 100000) }

func Benchmark_Int64_RadixLSD_100(b *testing.B)    { benchmarkInt64(b, Int64LSD, 100) }
func Benchmark_Int64_RadixLSD_1000(b *testing.B)   { benchmarkInt64(b, Int64LSD, 1000) }
func Benchmark_Int64_RadixLSD_10000(b *testing.B)  { benchmarkInt64(b, Int64LSD, 10000) }
func Benchmark_Int64_RadixLSD_100000(b *testing.B) { benchmarkInt64(b, Int64LSD, 100000) }

func Benchmark_Int64_StandardSort_100(b *testing.B)    { benchmarkInt64(b, int64_stdSort, 100) }
func Benchmark_Int64_StandardSort_1000(b *testing.B)   { benchmarkInt64(b, int64_stdSort, 1000) }
func Benchmark_Int64_StandardSort_10000(b *testing.B)  { benchmarkInt64(b, int64_stdSort, 10000) }
func Benchmark_Int64_StandardSort_100000(b *testing.B) { benchmarkInt64(b, int64_stdSort, 100000) }

func Benchmark_Int64_Insertion_100(b *testing.B) { benchmarkInt64(b, int64_insertion, 100) }

func benchmarkInt64(b *testing.B, sorter func([]int64), size int) {
	ys := make([][]int64, b.N)
	for n := range ys {
		ys[n] = int64_pop(size)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		sorter(ys[n])
	}
}

func int64_stdSort(xs []int64) {
	sort.Sort(byInt64(xs))
}

func int64_pop(size int) []int64 {
	xs := make([]int64, size)
	for i := range xs {
		xs[i] = int64(g.next())
	}
	return xs
}

type byInt64 []int64

func (xs byInt64) Len() int           { return len(xs) }
func (xs byInt64) Swap(i, j int)      { xs[i], xs[j] = xs[j], xs[i] }
func (xs byInt64) Less(i, j int) bool { return xs[i] < xs[j] }
