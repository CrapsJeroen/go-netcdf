// Copyright 2014 The Go-NetCDF Authors. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package netcdf

// #include <stddef.h>
import "C"

// cSizes copies s into a size_t array; size_t is 32-bit on 32-bit targets.
func cSizes(s []uint64) *C.size_t {
	if len(s) == 0 {
		return nil
	}
	c := make([]C.size_t, len(s))
	for i, v := range s {
		c[i] = C.size_t(v)
	}
	return &c[0]
}

// cPtrdiffs copies s into a ptrdiff_t array; ptrdiff_t is 32-bit on 32-bit targets.
func cPtrdiffs(s []int64) *C.ptrdiff_t {
	if len(s) == 0 {
		return nil
	}
	c := make([]C.ptrdiff_t, len(s))
	for i, v := range s {
		c[i] = C.ptrdiff_t(v)
	}
	return &c[0]
}
