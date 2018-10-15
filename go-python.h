#include "Python.h"
#include "frameobject.h"
#include "marshal.h"

/* stdlib */
#include <stdlib.h>
#include <string.h>

/* go-python */
#define _gopy_max_varargs 8 /* maximum number of varargs accepted by go-python */

PyObject*
_gopy_PyObject_CallFunction(PyObject *o, int len, char* types, void *args);