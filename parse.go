package gopython

//#cgo pkg-config: python-3.6
//#include "go-python.h"
import "C"

import (
	"fmt"
	"unsafe"
)

// PyObject* PyString_FromString(const char *v)
// Return value: New reference.
// Return a new string object with a copy of the string v as value on success, and NULL on failure. The parameter v must not be NULL; it will not be checked.
func PyString_FromString(v string) *PyObject {
	cstr := C.CString(v)
	defer C.free(unsafe.Pointer(cstr))
	return togo(C.PyBytes_FromString(cstr))
}

// char* PyString_AsString(PyObject *string)
// Return a NUL-terminated representation of the contents of string. The pointer refers to the internal buffer of string, not a copy. The data must not be modified in any way, unless the string was just created using PyString_FromStringAndSize(NULL, size). It must not be deallocated. If string is a Unicode object, this function computes the default encoding of string and operates on that. If string is not a string object at all, PyString_AsString() returns NULL and raises TypeError.
func PyString_AsString(self *PyObject) string {
	c_str := C.PyBytes_AsString(self.ptr)
	// we dont own c_str...
	//defer C.free(unsafe.Pointer(c_str))
	return C.GoString(c_str)
}

// long PyInt_AsLong(PyObject *io)
// Will first attempt to cast the object to a PyIntObject, if it is not already one, and then return its value. If there is an error, -1 is returned, and the caller should check PyErr_Occurred() to find out whether there was an error, or whether the value just happened to be -1.
func PyInt_AsLong(self *PyObject) int {
	return int(C.PyLong_AsLong(topy(self)))
}

// PyObject* PyInt_FromLong(long ival)
// Return value: New reference.
// Create a new integer object with a value of ival.
//
// The current implementation keeps an array of integer objects for all integers between -5 and 256, when you create an int in that range you actually just get back a reference to the existing object. So it should be possible to change the value of 1. I suspect the behaviour of Python in this case is undefined. :-)
func PyInt_FromLong(val int) *PyObject {
	return togo(C.PyLong_FromLong(C.long(val)))
}

// pyfmt returns the python format string for a given go value
func pyfmt(v interface{}) (unsafe.Pointer, string) {
	switch v := v.(type) {
	case bool:
		return unsafe.Pointer(&v), "b"

		// 	case byte:
		// 		return unsafe.Pointer(&v), "b"

	case int8:
		return unsafe.Pointer(&v), "b"

	case int16:
		return unsafe.Pointer(&v), "h"

	case int32:
		return unsafe.Pointer(&v), "i"

	case int64:
		return unsafe.Pointer(&v), "k"

	case int:
		switch unsafe.Sizeof(int(0)) {
		case 4:
			return unsafe.Pointer(&v), "i"
		case 8:
			return unsafe.Pointer(&v), "k"
		}

	case uint8:
		return unsafe.Pointer(&v), "B"

	case uint16:
		return unsafe.Pointer(&v), "H"

	case uint32:
		return unsafe.Pointer(&v), "I"

	case uint64:
		return unsafe.Pointer(&v), "K"

	case uint:
		switch unsafe.Sizeof(uint(0)) {
		case 4:
			return unsafe.Pointer(&v), "I"
		case 8:
			return unsafe.Pointer(&v), "K"
		}

	case float32:
		return unsafe.Pointer(&v), "f"

	case float64:
		return unsafe.Pointer(&v), "d"

	case complex64:
		return unsafe.Pointer(&v), "D"

	case complex128:
		return unsafe.Pointer(&v), "D"

	case string:
		cstr := C.CString(v)
		return unsafe.Pointer(cstr), "s"

	case *PyObject:
		return unsafe.Pointer(v.topy()), "O"

	}

	panic(fmt.Errorf("python: unknown type (%T)", v))
}
