package gopython

import (
	"errors"
	"strings"
)

var PyStr = PyString_FromString
var GoStr = PyString_AsString
var GoInt = PyInt_AsLong

// InsertExtraPackagePath 添加额外的包引用路径
func InsertExtraPackagePath(dir string) (*PyObject, error) {
	sysModule := PyImport_ImportModule("sys")
	path := sysModule.GetAttrString("path")
	if path != nil {
		if str := GoStr(path.Repr()); !strings.Contains(str, dir) {
			if err := PyList_Insert(path, 0, PyStr(dir)); err != nil {
				return nil, err
			}
		}

		return path, nil
	}

	return nil, errors.New("未导入指定模块路径")
}

// ImportModule 从指定的文件夹中导入包
func ImportModule(name string) (*PyObject, error) {
	if obj := PyImport_ImportModule(name); obj != nil {
		return obj, nil
	}

	return nil, errors.New("未导入指定模块")
}

// CallFunc 调用 Python 中的方法
// 指定模块名称，方法名称，排列参数
func CallFunc(modulename string, funcname string, args ...interface{}) (*PyObject, error) {
	module, err := getModule(modulename)
	if err == nil {
		if len(args) > 0 {
			funcArgs := PyTuple_New(len(args))
			for i := 0; i < len(args); i++ {
				switch args[i].(type) {
				case int:
					PyTuple_SetItem(funcArgs, i, PyInt_FromLong(args[i].(int)))
				case string:
					PyTuple_SetItem(funcArgs, i, PyStr(args[i].(string)))
				}
			}

			res := module.GetAttrString(funcname).Call(funcArgs, Py_None)
			return res, nil
		}

		res := module.GetAttrString(funcname).CallFunction()
		return res, nil
	} else {
		return nil, err
	}
}

// getModule 获得导入模块的引用
// TODO:使用其他诸如缓存方法获取模块，取代重新导入
func getModule(modulename string) (*PyObject, error) {
	return ImportModule(modulename)
}
