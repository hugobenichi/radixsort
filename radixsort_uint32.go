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
	int32_least_significant_digit(*(*[]int32)(unsafe.Pointer(&xs)), 0)
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
