package radix

// xs64s is a xorshift64 star pseudo-random-number-generator to use in tests.
type xs64s uint64

func prng() xs64s { return 1 }

func (r *xs64s) next() uint64 {
	u := *r

	u ^= u >> 12
	u ^= u << 25
	u ^= u >> 27

	*r = u

	return uint64(u) * 2685821657736338717
}
