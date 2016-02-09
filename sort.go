// package radix provides radix sort for int32 by most significant bit and least
// significant bit.
package radix

import (
	"fmt"
)

func integr(cs, is *[256]uint32) {
	a := uint32(0)
	for i := 0; i < 256; i++ {
		c := cs[i]
		is[i] = a
		a += c
	}
}

func count(xs []int32, cs *[256]uint32, offset int32, shift uint) {
	for _, x := range xs {
		r := (offset + (x >> shift)) & 0xFF
		cs[r]++
	}
}

func swap(xs, ys []int32, is *[256]uint32, offset int32, shift uint) {
	for _, x := range xs {
		r := (offset + (x >> shift)) & 0xFF
		ys[is[r]] = x
		is[r]++
	}
	copy(xs, ys)
}

func radSortAt(xs, ys []int32, cs, is *[256]uint32, offset int32, shift uint) {
	count(xs, cs, offset, shift)
	integr(cs, is)
	swap(xs, ys, is, offset, shift)
}

func Int32MSB(xs []int32) {
	if len(xs) <= 64 {
		int32_insertion(xs)
		return
	}

	var (
		cs, is [256]uint32
		ys     = make([]int32, len(xs))
	)
	radSortAt(xs, ys, &cs, &is, 1<<7, 24)

	// partially sort every radix bucket by with secondary radix sort when count is too high
	var lo, hi uint32
	for i := 0; i < 256; i++ {
		c := cs[i]
		if c == 0 {
			continue
		}
		hi = lo + c
		zs := xs[lo:hi]
		if c > 20000 {
			var ds [256]uint32
			radSortAt(zs, ys, &ds, &ds, 0, 8)
		}
		if c > 100 {
			var ds [256]uint32
			radSortAt(zs, ys, &ds, &ds, 0, 16)
		}
		lo = hi
	}

	int32_insertion(xs) // ~linear runtime when globally sorted, locally not-sorted
}

// Int32LSB sorts in place the given array of int32 numbers using least
// significant digit radix sort. It uses additional swap space equal to the
// given array length. When the length of the given array is equal or less than
// 64, insertion sort is used instead.
func Int32LSB(xs []int32) {
	if len(xs) <= 64 {
		int32_insertion(xs)
		return
	}

	var css [4][256]uint32 // should be living on the stack

	// count all radix keys
	for _, x := range xs {
		var (
			a = x & 0xFF
			b = (x >> 8) & 0xFF
			c = (x >> 16) & 0xFF
			d = ((1 << 7) + (x >> 24)) & 0xFF // translate by +128 for signed order
		)
		css[0][a]++
		css[1][b]++
		css[2][c]++
		css[3][d]++
	}

	// aggregate radix counts to radix offsets
	for i := range css {
		cs := &css[i]
		a := uint32(0)
		for j := 0; j < 256; j++ {
			c := cs[j]
			cs[j] = a
			a += c
		}
	}

	var (
		ys = make([]int32, len(xs)) // temp array for swapping elements
		ss = [4]uint{0, 8, 16, 24}
		os = [4]int32{0, 0, 0, 1 << 7}
	)
	for i := range css {
		var (
			cs     = css[i] // do not obtain cs from range expr
			shift  = ss[i]
			offset = os[i]
		)
		for _, x := range xs {
			r := (offset + (x >> shift)) & 0xFF
			j := cs[r]
			cs[r]++
			ys[j] = x
		}
		xs, ys = ys, xs // even number of swap
	}
}

func int32_insertion(xs []int32) {
	for i := 1; i < len(xs); i++ {
		j, x := i, xs[i]
		//a := 0
		for j > 0 && xs[j-1] > x {
			xs[j] = xs[j-1]
			j--
			//a++
		}
		//fmt.Println(a)
		xs[j] = x
	}
}

var _not_used = fmt.Println
