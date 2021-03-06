package radixsort

import (
	"unsafe"
)

// Radix sort for int64. Int64 delegates to most significant digit radix sort.
func Int64(xs []int64) { Int64MSD(xs) }

// Radix sort for uint64. Uint64 delegates to most significant digit radix sort.
func Uint64(xs []uint64) { Uint64MSD(xs) }

// Most significant digit radix sort for int64.
func Int64MSD(xs []int64) {
	if len(xs) <= 64 {
		int64_insertion(xs)
		return
	}
	var (
		temp = make([]int64, len(xs))
		is   [256]uint32
	)
	int64_most_significant_digit(xs, temp, &is, 1<<7, 56)
}

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
	int64_most_significant_digit(*(*[]int64)(unsafe.Pointer(&xs)), temp, &is, 0, 56)
}

// Least significant digit radix sort for int64.
func Int64LSD(xs []int64) {
	if len(xs) <= 64 {
		int64_insertion(xs)
		return
	}
	int64_least_significant_digit(xs, 1<<7)
}

// Least significant digit radix sort for uint64.
func Uint64LSD(xs []uint64) {
	if len(xs) <= 64 {
		uint64_insertion(xs)
		return
	}
	int64_least_significant_digit(*(*[]int64)(unsafe.Pointer(&xs)), 0)
}

func int64_most_significant_digit(xs, temp []int64, is *[256]uint32, offset int64, shift uint) {
	var cs [256]uint32
	for _, x := range xs {
		r := (offset + (x >> shift)) & 0xFF
		cs[r]++
	}
	a := uint32(0)
	for i := 0; i < 256; i++ {
		is[i] = a
		a += cs[i]
	}
	for _, x := range xs {
		r := (offset + (x >> shift)) & 0xFF
		temp[is[r]] = x
		is[r]++
	}
	copy(xs, temp)

	if shift == 0 { // that was the last radix digit
		return
	}

	var lo uint32
	for i := 0; i < 256; i++ {
		var (
			c  = cs[i]
			hi = lo + c
			zs = xs[lo:hi]
		)
		lo = hi

		switch {
		case c < 2: // already sorted
		case c <= 100:
			int64_insertion(zs) // ~linear runtime when globally sorted, locally not-sorted
		default:
			int64_most_significant_digit(zs, temp, is, 0, shift-8)
		}
	}
}

func int64_least_significant_digit(xs []int64, offsetMSD int64) {
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
			h = (offsetMSD + (x >> 56)) & 0xFF // translate by +128 for signed order
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
		ys = make([]int64, len(xs)) // temp array for swapping elements
		ss = [8]uint{0, 8, 16, 24, 32, 40, 48, 56}
		os = [8]int64{0, 0, 0, 0, 0, 0, 0, offsetMSD}
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

func int64_insertion(xs []int64) {
	for i := 1; i < len(xs); i++ {
		j, x := i, xs[i]
		for j > 0 && xs[j-1] > x {
			xs[j] = xs[j-1]
			j--
		}
		xs[j] = x
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
