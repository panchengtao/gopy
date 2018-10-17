package gopython

//#cgo pkg-config: python-3.6
//#include "go-python.h"
import "C"

import (
	"fmt"
	"strings"
	"unsafe"
)

// The Python None object, denoting lack of value. This object has no methods.
// It needs to be treated just like any other object with respect to reference
// counts.
var Py_None = toGoPyObject(C.Py_None)

// PyThreadState layer
type PyThreadState struct {
	ptr *C.PyThreadState
}

// PyObject layer
type PyObject struct {
	ptr *C.PyObject
}

// PyObject* PyObject_GetAttrString(PyObject *o, const char *attr_name)
// Return value: New reference.
// Retrieve an attribute named attr_name from object o. Returns the attribute value on success, or NULL on failure. This is the equivalent of the Python expression o.attr_name.
func (self *PyObject) GetAttrString(attr_name string) *PyObject {
	c_attr_name := C.CString(attr_name)
	defer C.free(unsafe.Pointer(c_attr_name))
	return toGoPyObject(C.PyObject_GetAttrString(self.ptr, c_attr_name))
}

// PyObject* PyObject_Repr(PyObject *o)
// Return value: New reference.
// Compute a string representation of object o. Returns the string representation on success, NULL on failure. This is the equivalent of the Python expression repr(o). Called by the repr() built-in function and by reverse quotes.
func (self *PyObject) Repr() *PyObject {
	return toGoPyObject(C.PyObject_Str(self.ptr))
}

// PyObject* PyObject_Call(PyObject *callable_object, PyObject *args, PyObject *kw)
// Return value: New reference.
// Call a callable Python object callable_object, with arguments given by the tuple args, and named arguments given by the dictionary kw. If no named arguments are needed, kw may be NULL. args must not be NULL, use an empty tuple if no arguments are needed. Returns the result of the call on success, or NULL on failure. This is the equivalent of the Python expression apply(callable_object, args, kw) or callable_object(*args, **kw).
func (self *PyObject) Call(args, kw *PyObject) *PyObject {
	return toGoPyObject(C.PyObject_Call(self.ptr, args.ptr, kw.ptr))
}

// PyObject* PyObject_CallFunction(PyObject *callable, char *format, ...)
// Return value: New reference.
// Call a callable Python object callable, with a variable number of C arguments. The C arguments are described using a Py_BuildValue() style format string. The format may be NULL, indicating that no arguments are provided. Returns the result of the call on success, or NULL on failure. This is the equivalent of the Python expression apply(callable, args) or callable(*args). Note that if you only pass PyObject * args, PyObject_CallFunctionObjArgs() is a faster alternative.
func (self *PyObject) CallFunction(args ...interface{}) *PyObject {
	if len(args) > int(C._gopy_max_varargs) {
		panic(fmt.Errorf(
			"gopy: maximum number of varargs (%d) exceeded (%d)",
			int(C._gopy_max_varargs),
			len(args),
		))
	}

	types := make([]string, 0, len(args))
	cargs := make([]unsafe.Pointer, 0, len(args))

	for _, arg := range args {
		ptr, typ := pyfmt(arg)
		types = append(types, typ)
		cargs = append(cargs, ptr)
		if typ == "s" {
			defer func(ptr unsafe.Pointer) {
				C.free(ptr)
			}(ptr)
		}
	}

	if len(args) <= 0 {
		o := C._gopy_PyObject_CallFunction(self.ptr, 0, nil, nil)
		return toGoPyObject(o)
	}

	pyfmt := C.CString(strings.Join(types, ""))
	defer C.free(unsafe.Pointer(pyfmt))
	o := C._gopy_PyObject_CallFunction(
		self.ptr,
		C.int(len(args)),
		pyfmt,
		unsafe.Pointer(&cargs[0]),
	)

	return toGoPyObject(o)

}

func (self *PyObject) topy() *C.PyObject {
	return self.ptr
}
