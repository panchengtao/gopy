#include "go-python.h"

PyObject*
_gopy_PyObject_CallFunction(PyObject *o, int len, char* pyfmt, void *cargs) {
	void ** args = (void**)cargs;

	if (len > _gopy_max_varargs) {
			PyErr_Format(
					PyExc_RuntimeError,
					"python: maximum number of varargs (%d) exceeded (%d)",
					_gopy_max_varargs,
					len
			);
			return NULL;
	}

	switch (len) {
		case 0:
			return PyObject_CallFunction(o, pyfmt);

		case 1:
			return PyObject_CallFunction(o, pyfmt, args[0]);

		case 2:
			return PyObject_CallFunction(o, pyfmt, args[0], args[1]);

		case 3:
			return PyObject_CallFunction(o, pyfmt, args[0], args[1], args[2]);

		case 4:
			return PyObject_CallFunction(o, pyfmt, args[0], args[1], args[2], args[3]);

		case 5:
			return PyObject_CallFunction(o, pyfmt, args[0], args[1], args[2], args[3], args[4]);

		case 6:
			return PyObject_CallFunction(o, pyfmt, args[0], args[1], args[2], args[3], args[4], args[5]);

		case 7:
			return PyObject_CallFunction(o, pyfmt, args[0], args[1], args[2], args[3], args[4], args[5], args[6]);

		case 8:
			return PyObject_CallFunction(o, pyfmt, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7]);

		default:
			PyErr_Format(PyExc_RuntimeError, "python: invalid number of arguments (%d)", len);
			return NULL;

	}

	return NULL;
}