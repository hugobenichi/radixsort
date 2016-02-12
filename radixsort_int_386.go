// +build 386
package radixsort

import (
	"unsafe"
)

// Radix sort for int. Int delegates to int32 least significant digit radix sort.
// Only works on 32bits architectures.
func Int(xs []int) {
	Int32LSD(*(*[]int32)(unsafe.Pointer(&xs)))
}

// Radix sort for uint. Uint delegates to uint32 least significant digit radix sort.
// Only works on 32bits architectures.
func Uint(xs []int) {
	Uint32LSD(*(*[]uint32)(unsafe.Pointer(&xs)))
}
