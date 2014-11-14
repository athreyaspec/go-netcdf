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

// PutUshort writes data as the entire data for variable v.
func (v Var) PutUshort(data []uint16) error {
	if err := v.okData(NC_USHORT, len(data)); err != nil {
		return err
	}
	return newError(C.nc_put_var_ushort(C.int(v.f), C.int(v.id), (*C.ushort)(unsafe.Pointer(&data[0]))))
}

// GetUshort reads the entire variable v into data, which must have enough
// space for all the values (i.e. len(data) must be at least v.Len()).
func (v Var) GetUshort(data []uint16) error {
	if err := v.okData(NC_USHORT, len(data)); err != nil {
		return err
	}
	return newError(C.nc_get_var_ushort(C.int(v.f), C.int(v.id), (*C.ushort)(unsafe.Pointer(&data[0]))))
}

// PutUshort sets the value of attribute a to val.
func (a Attr) PutUshort(val []uint16) error {
	// TODO: check Type is NC_DOUBLE and len(val) is corrent
	cname := C.CString(a.name)
	defer C.free(unsafe.Pointer(cname))
	return newError(C.nc_put_att_ushort(C.int(a.v.f), C.int(a.v.id), cname,
		C.nc_type(NC_USHORT), C.size_t(len(val)), (*C.ushort)(unsafe.Pointer(&val[0]))))
}

// GetUshort returns the attribute value.
func (a Attr) GetUshort() (val []uint16, err error) {
	// TODO: check Type is NC_DOUBLE
	n, err := a.Len()
	if err != nil {
		return nil, err
	}
	cname := C.CString(a.name)
	defer C.free(unsafe.Pointer(cname))
	val = make([]uint16, n)
	err = newError(C.nc_get_att_ushort(C.int(a.v.f), C.int(a.v.id), cname,
		(*C.ushort)(unsafe.Pointer(&val[0]))))
	return
}