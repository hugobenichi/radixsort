package radix

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	const (
		rep = 5
		ln  = 200000 // crossing MSB/LSB @ ~200000
	)

	var (
		xs  = make([]int32, ln)
		tot = int64(0)
	)

	for i := 0; i < rep; i++ {
		pop(xs)

		s := time.Now()
		//radix.Int32MSB(xs)
		//radix.Int32LSB(xs)
		sort.Sort(byInt32(xs))
		e := time.Now()

		tot += e.UnixNano() - s.UnixNano()
	}

	// TODO: replace with sort pkg
	fmt.Println(Ints32AreSorted(xs), tot/rep)
}

func pop(xs []int32) {
	for i := range xs {
		xs[i] = int32(rand.Int())
		if rand.Int()&1 == 1 {
			xs[i] = ^xs[i]
		}
	}
}

func print(xs []int32) {
	for _, x := range xs {
		fmt.Println(x)
	}
}

func Ints32AreSorted(xs []int32) bool {
	if len(xs) < 2 {
		return true
	}
	y := xs[0]
	for _, x := range xs {
		if y > x {
			return false
		}
		y = x
	}
	return true
}

var (
	bla = sort.Ints
)

type byInt32 []int32

func (xs byInt32) Len() int           { return len(xs) }
func (xs byInt32) Swap(i, j int)      { xs[i], xs[j] = xs[j], xs[i] }
func (xs byInt32) Less(i, j int) bool { return xs[i] < xs[j] }
