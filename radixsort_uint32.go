package radixsort

import (
	"unsafe"
)

// Radix sort for uint32. Uint32 delegates to least significant digit radix sort.
func Uint32(xs []uint32) { Uint32LSD(xs) }

// Most significant digit radix sort for uint32.
func Uint32MSD(xs []uint32) {
	if len(xs) <= 64 {
		uint32_insertion(xs)
		return
	}
	var (
		temp = make([]int32, len(xs))
		is   [256]uint32
	)
	int32_sortAtRadix_rec(*(*[]int32)(unsafe.Pointer(&xs)), temp, &is, 0, 24)
}

// Least significant digit radix sort for uint32.
func Uint32LSD(xs []uint32) {
	if len(xs) <= 64 {
		uint32_insertion(xs)
		return
	}

	var css [4][256]uint32 // should be living on the stack

	// count all radix keys
	for _, x := range xs {
		var (
			a = x & 0xFF
			b = (x >> 8) & 0xFF
			c = (x >> 16) & 0xFF
			d = (x >> 24) & 0xFF
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
		ys = make([]uint32, len(xs)) // temp array for swapping elements
		ss = [4]uint{0, 8, 16, 24}
	)
	for i := range css {
		var (
			cs    = css[i] // do not obtain cs from range expr
			shift = ss[i]
		)
		for _, x := range xs {
			r := (x >> shift) & 0xFF
			j := cs[r]
			cs[r]++
			ys[j] = x
		}
		xs, ys = ys, xs // even number of swap
	}
}

func uint32_insertion(xs []uint32) {
	for i := 1; i < len(xs); i++ {
		j, x := i, xs[i]
		for j > 0 && xs[j-1] > x {
			xs[j] = xs[j-1]
			j--
		}
		xs[j] = x
	}
}
