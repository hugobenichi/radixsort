// package radix provides radix sort for int64 by most significant bit and least
// significant bit.
package radix

func int64_sortAtRadix(xs, ys []int64, cs, is *[256]uint32, offset int64, shift uint) {
	for _, x := range xs {
		r := (offset + (x >> shift)) & 0xFF
		cs[r]++
	}
	a := uint32(0)
	for i := 0; i < 256; i++ {
		c := cs[i]
		is[i] = a
		a += c
	}
	for _, x := range xs {
		r := (offset + (x >> shift)) & 0xFF
		ys[is[r]] = x
		is[r]++
	}
	copy(xs, ys)
}

func Int64MSB(xs []int64) {
	if len(xs) <= 64 {
		int64_insertion(xs)
		return
	}

	var (
		cs, is [256]uint32
		ys     = make([]int64, len(xs))
	)
	int64_sortAtRadix(xs, ys, &cs, &is, 1<<7, 56)

	// partially sort every radix bucket by with secondary radix sort when count is too high
	var lo uint32
	for i := 0; i < 256; i++ {
		var (
			c  = cs[i]
			hi = lo + c
			zs = xs[lo:hi]
		)
		lo = hi

		if c < 2 { // already sorted
			continue
		}
		if c > 20000 {
			for i := 0; i < 256; i++ {
				is[i] = 0
			}
			int64_sortAtRadix(zs, ys, &is, &is, 0, 40)
		}
		if c > 100 {
			for i := 0; i < 256; i++ {
				is[i] = 0
			}
			int64_sortAtRadix(zs, ys, &is, &is, 0, 48)
		}
		int64_insertion(zs) // ~linear runtime when globally sorted, locally not-sorted
	}
}

func int64_sortAtRadix_rec(xs, temp []int64, is *[256]uint32, offset int64, shift uint) {
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

	if shift == 0 { // that was the last radix bucket
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
			int64_insertion(zs)
		default:
			int64_sortAtRadix_rec(zs, temp, is, 0, shift-8)
		}
	}
}

func Int64MSB_alt(xs []int64) {
	if len(xs) <= 64 {
		int64_insertion(xs)
		return
	}
	var (
		temp = make([]int64, len(xs))
		is   [256]uint32
	)
	int64_sortAtRadix_rec(xs, temp, &is, 1<<7, 56)
}

// Int64LSB sorts in place the given array of int64 numbers using least
// significant digit radix sort. It uses additional swap space equal to the
// given array length. When the length of the given array is equal or less than
// 64, insertion sort is used instead.
func Int64LSB(xs []int64) {
	if len(xs) <= 64 {
		int64_insertion(xs)
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
			h = ((1 << 7) + (x >> 56)) & 0xFF // translate by +128 for signed order
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
		os = [8]int64{0, 0, 0, 0, 0, 0, 0, 1 << 7}
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
