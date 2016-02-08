// TODO: descr
package radix

import (
	"fmt"
)

func integr(cs, is *[256]uint32) {
	a := uint32(0)
	for i, c := range cs {
		is[i] = a
		a += c
	}
}

func count(xs []int32, cs *[256]uint32, rad func(int32) int32) {
	for _, x := range xs {
		cs[rad(x)]++
	}
}

func radSwap(xs, ys []int32, is *[256]uint32, rad func(int32) int32) {
	for _, x := range xs {
		r := rad(x)
		ys[is[r]] = x
		is[r]++
	}
	copy(xs, ys)
}

func signedrad(x int32) int32 { // int is 64 bits
	// translate by +128 for signed radix order
	return ((1 << 7) + (x >> 24)) & 0xFF
}

func rad(x int32) int32 {
	return (x >> 24) & 0xff
}

func radAt(shift uint) func(int32) int32 {
	return func(x int32) int32 {
		return (x >> shift) & 0xff
	}
}

func radSortAt(xs, ys []int32, cs, is *[256]uint32, rad func(int32) int32) {
	count(xs, cs, rad)
	integr(cs, is)
	radSwap(xs, ys, is, rad)
}

func Int32MSB(xs []int32) {
	var (
		cs, is [256]uint32
		ys     = make([]int32, len(xs))
	)

	for _, x := range xs {
		cs[signedrad(x)]++
	}
	a := uint32(0)
	for i, c := range cs {
		is[i] = a
		a += c
	}
	for _, x := range xs {
		r := signedrad(x)
		ys[is[r]] = x
		is[r]++
	}
	for i := range xs {
		xs[i] = ys[i]
	}

	// partially sort every radix bucket by with secondary radix sort when count is too high
	var lo, hi uint32
	for _, c := range cs {
		if c == 0 {
			continue
		}
		hi = lo + c
		zs := xs[lo:hi]
		if c > 20000 {
			for i := range is {
				is[i] = 0
			}
			for _, x := range zs {
				r := (x >> 8) & 0xFF
				is[r]++
			}
			a = uint32(0)
			for i, c := range is {
				is[i] = a
				a += c
			}
			for _, x := range zs {
				r := (x >> 8) & 0xFF
				ys[is[r]] = x
				is[r]++
			}
			for i := range zs {
				zs[i] = ys[i]
			}
		}
		if c > 255 {
			for i := range is {
				is[i] = 0
			}
			for _, x := range zs {
				r := (x >> 16) & 0xFF
				is[r]++
			}
			a = uint32(0)
			for i, c := range is {
				is[i] = a
				a += c
			}
			for _, x := range zs {
				r := (x >> 16) & 0xFF
				ys[is[r]] = x
				is[r]++
			}
			for i := range zs {
				zs[i] = ys[i]
			}
		}
		lo = hi
	}
	insertion(xs) // efficient when partially sorted
}

func Int32LSB(xs []int32) {
	var css [4][256]uint32
	for _, x := range xs {
		var (
			a = x & 0xFF
			b = (x >> 8) & 0xFF
			c = (x >> 16) & 0xFF
			d = ((1 << 7) + (x >> 24)) & 0xFF
		)
		css[0][a]++
		css[1][b]++
		css[2][c]++
		css[3][d]++
	}

	ys := make([]int32, len(xs))
	shift := uint(0)
	offset := int32(0)
	for i := range css {
		cs := css[i]
		integr(&cs, &cs)
		for _, x := range xs {
			r := (offset + (x >> shift)) & 0xFF
			j := cs[r]
			cs[r]++
			ys[j] = x
		}
		offset = 0
		xs, ys = ys, xs
		shift += 8
		if shift == 24 {
			offset = 1 << 7
		}
	}
}

func insertion(xs []int32) {
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

var gaps = []int{57, 23, 10, 4, 1}

func shell(xs []int32) {
	for _, g := range gaps {
		for i := g; i < len(xs); i++ {
			x := xs[i]
			j := i
			for j >= g && xs[j-g] > x {
				xs[j] = xs[j-g]
				j -= g
			}
			xs[j] = x
		}
	}
}

var blabla = fmt.Println
