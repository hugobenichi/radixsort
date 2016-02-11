package radixsort

// xs64s is a xorshift64 star pseudo-random-number-generator to use in tests.
type xs64s uint64

func (r *xs64s) next() uint64 {
	u := *r

	u ^= u >> 12
	u ^= u << 25
	u ^= u >> 27

	*r = u

	return uint64(u) * 2685821657736338717
}

var g = xs64s(1)
