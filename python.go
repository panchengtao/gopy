package gopython

//#cgo pkg-config: python-3.6
//#include "go-python.h"
import "C"

import (
	"fmt"
	"github.com/pkg/errors"
	"unsafe"
)

// Initialize initializes the python interpreter and its GIL
func Initialize() error {
	// make sure the python interpreter has been initialized
	if C.Py_IsInitialized() == 0 {
		C.Py_Initialize()
	}
	if C.Py_IsInitialized() == 0 {
		return fmt.Errorf("python: could not initialize the python interpreter")
	}

	// make sure the GIL is correctly initialized
	if C.PyEval_ThreadsInitialized() == 0 {
		C.PyEval_InitThreads()
	}
	if C.PyEval_ThreadsInitialized() == 0 {
		return fmt.Errorf("python: could not initialize the GIL")
	}

	return nil
}

// Finalize shutdowns the python interpreter
func Finalize() error {
	C.Py_Finalize()
	return nil
}

// PyObject* PyImport_ImportModule(const char *name)
// Return value: New reference.
// This is a simplified interface to PyImport_ImportModuleEx() below, leaving the globals and locals arguments set to NULL and level set to 0. When the name argument contains a dot (when it specifies a submodule of a package), the fromlist argument is set to the list ['*'] so that the return value is the named module rather than the top-level package containing it as would otherwise be the case. (Unfortunately, this has an additional side effect when name in fact specifies a subpackage instead of a submodule: the submodules specified in the package’s __all__ variable are loaded.) Return a new reference to the imported module, or NULL with an exception set on failure. Before Python 2.4, the module may still be created in the failure case — examine sys.modules to find out. Starting with Python 2.4, a failing import of a module no longer leaves the module in sys.modules.
//
// Changed in version 2.4: Failing imports remove incomplete module objects.
//
// Changed in version 2.6: Always uses absolute imports.
func PyImport_ImportModule(name string) *PyObject {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	return togo(C.PyImport_ImportModule(c_name))
}

// PyObject* PyTuple_New(Py_ssize_t len)
// Return value: New reference.
// Return a new tuple object of size len, or NULL on failure.
//
// Changed in version 2.5: This function used an int type for len. This might require changes in your code for properly supporting 64-bit systems.
func PyTuple_New(sz int) *PyObject {
	return togo(C.PyTuple_New(C.Py_ssize_t(sz)))
}

// int PyTuple_SetItem(PyObject *p, Py_ssize_t pos, PyObject *o)
// Insert a reference to object o at position pos of the tuple pointed to by p. Return 0 on success.
//
// Note This function “steals” a reference to o.
// Changed in version 2.5: This function used an int type for pos. This might require changes in your code for properly supporting 64-bit systems.
func PyTuple_SetItem(self *PyObject, pos int, o *PyObject) error {
	err := C.PyTuple_SetItem(topy(self), C.Py_ssize_t(pos), topy(o))
	if err == 0 {
		return nil
	}

	return errors.New(fmt.Sprintf("error in C-Python (rc=%d)", int(err)))
}

// int PyRun_SimpleString(const char *command)
// This is a simplified interface to PyRun_SimpleStringFlags() below, leaving the PyCompilerFlags* argument set to NULL.
func PyRun_SimpleString(command string) int {
	c_cmd := C.CString(command)
	defer C.free(unsafe.Pointer(c_cmd))
	return int(C._gopy_PyRun_SimpleString(c_cmd))
}