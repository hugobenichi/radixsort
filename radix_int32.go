// package radix provides radix sort for int32 by most significant bit and least
// significant bit.
package radix

func Int32MSB(xs []int32) {
	if len(xs) <= 64 {
		int32_insertion(xs)
		return
	}
	var (
		temp = make([]int32, len(xs))
		is   [256]uint32
	)
	int32_sortAtRadix_rec(xs, temp, &is, 1<<7, 24)
}

func int32_sortAtRadix_rec(xs, temp []int32, is *[256]uint32, offset int32, shift uint) {
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
			int32_insertion(zs) // ~linear runtime when globally sorted, locally not-sorted
		default:
			int32_sortAtRadix_rec(zs, temp, is, 0, shift-8)
		}
	}
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
		for j > 0 && xs[j-1] > x {
			xs[j] = xs[j-1]
			j--
		}
		xs[j] = x
	}
}
