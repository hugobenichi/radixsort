package radixsort

import (
	"unsafe"
)

// Radix sort for int. Int delegates to int64 most significant digit radix sort.
// Only works on 64bits architectures.
func Int(xs []int) {
	Int64MSD(*(*[]int64)(unsafe.Pointer(&xs)))
}

// Radix sort for uint. Uint delegates to uint64 most significant digit radix sort.
// Only works on 64bits architectures.
func Uint(xs []uint) {
	Uint64MSD(*(*[]uint64)(unsafe.Pointer(&xs)))
}
