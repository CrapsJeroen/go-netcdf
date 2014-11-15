// Copyright 2014 The Go-NetCDF Authors. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// These files are autogenerated from nc_double.go using generate.sh

package netcdf

import (
	"unsafe"
)

// #include <stdlib.h>
// #include <netcdf.h>
import "C"

// WriteChar writes data as the entire data for variable v.
func (v Var) WriteChar(data []byte) error {
	if err := okData(v, NC_CHAR, len(data)); err != nil {
		return err
	}
	return newError(C.nc_put_var_text(C.int(v.f), C.int(v.id), (*C.char)(unsafe.Pointer(&data[0]))))
}

// ReadChar reads the entire variable v into data, which must have enough
// space for all the values (i.e. len(data) must be at least v.Len()).
func (v Var) ReadChar(data []byte) error {
	if err := okData(v, NC_CHAR, len(data)); err != nil {
		return err
	}
	return newError(C.nc_get_var_text(C.int(v.f), C.int(v.id), (*C.char)(unsafe.Pointer(&data[0]))))
}

// WriteChar sets the value of attribute a to val.
func (a Attr) WriteChar(val []byte) error {
	// We don't need okData here because netcdf library doesn't know
	// the length or type of the attribute yet.
	cname := C.CString(a.name)
	defer C.free(unsafe.Pointer(cname))
	return newError(C.nc_put_att_text(C.int(a.v.f), C.int(a.v.id), cname,
		C.size_t(len(val)), (*C.char)(unsafe.Pointer(&val[0]))))
}

// ReadChar reads the entire attribute value into val.
func (a Attr) ReadChar(val []byte) (err error) {
	if err := okData(a, NC_CHAR, len(val)); err != nil {
		return err
	}
	cname := C.CString(a.name)
	defer C.free(unsafe.Pointer(cname))
	err = newError(C.nc_get_att_text(C.int(a.v.f), C.int(a.v.id), cname,
		(*C.char)(unsafe.Pointer(&val[0]))))
	return
}

// CharReader is a interface that allows reading a sequence of values of fixed length.
type CharReader interface {
	Len() (n uint64, err error)
	ReadChar(val []byte) (err error)
}

// GetChar reads the entire data in r and returns it.
func GetChar(r CharReader) (data []byte, err error) {
	n, err := r.Len()
	if err != nil {
		return
	}
	data = make([]byte, n)
	err = r.ReadChar(data)
	return
}