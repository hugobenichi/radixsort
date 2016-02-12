package radixsort

import (
	"unsafe"
)

// Radix sort for uint64. Uint64 delegates to most significant digit radix sort.
func Uint64(xs []uint64) { Uint64MSD(xs) }

// Most significant digit radix sort for uint64.
func Uint64MSD(xs []uint64) {
	if len(xs) <= 64 {
		uint64_insertion(xs)
		return
	}
	var (
		temp = make([]int64, len(xs))
		is   [256]uint32
	)
	int64_sortAtRadix_rec(*(*[]int64)(unsafe.Pointer(&xs)), temp, &is, 0, 56)
}

// Least significant digit radix sort for uint64.
func Uint64LSD(xs []uint64) {
	if len(xs) <= 64 {
		uint64_insertion(xs)
		return
	}

	var css [8][256]uint32 // should be living on the stack

	// count all radix keys
	for _, x := range xs {
		var (
			a = x & 0xFF
			b = (x >> 8) & 0xFF
			c = (x >> 16) & 0xFF
			d = (x >> 24) & 0xFF
			e = (x >> 32) & 0xFF
			f = (x >> 40) & 0xFF
			g = (x >> 48) & 0xFF
			h = (x >> 56) & 0xFF // translate by +128 for signed order
		)
		css[0][a]++
		css[1][b]++
		css[2][c]++
		css[3][d]++
		css[4][e]++
		css[5][f]++
		css[6][g]++
		css[7][h]++
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
		ys = make([]uint64, len(xs)) // temp array for swapping elements
		ss = [8]uint{0, 8, 16, 24, 32, 40, 48, 56}
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

func uint64_insertion(xs []uint64) {
	for i := 1; i < len(xs); i++ {
		j, x := i, xs[i]
		for j > 0 && xs[j-1] > x {
			xs[j] = xs[j-1]
			j--
		}
		xs[j] = x
	}
}
