package gopython

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCompleteWorkFlow(t *testing.T) {
	testFile, err := os.OpenFile("./hello.py", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	assert.Nil(t, err)
	testFile.Write([]byte("def minus(x,y):\n"))
	testFile.Write([]byte("	   return x - y\n"))
	testFile.Write([]byte("def add(x,y):\n"))
	testFile.Write([]byte("    return x + y\n"))
	testFile.Write([]byte("def return_self(self):\n"))
	testFile.Write([]byte("    return self\n"))
	testFile.Write([]byte("def zero_arg():\n"))
	testFile.Write([]byte("    return 0\n"))
	testFile.Close()

	err = Initialize()
	assert.Nil(t, err)

	InsertPackagePath("./")

	ret, err := CallFunc("hello", "minus", 1, 1)
	callInt := GoInt(ret)
	assert.True(t, callInt == 0)
	assert.Nil(t, err)

	ret, err = CallFunc("hello", "add", 1, 1)
	callInt = GoInt(ret)
	assert.True(t, callInt == 2)
	assert.Nil(t, err)

	self1, err := CallFunc("hello", "return_self", 1)
	callSelf1 := GoInt(self1)
	assert.True(t, callSelf1 == 1)
	assert.Nil(t, err)

	self2, err := CallFunc("hello", "return_self", "self")
	callSelf2 := GoStr(self2)
	assert.True(t, callSelf2 == "self")
	assert.Nil(t, err)

	zero, err := CallFunc("hello", "zero_arg")
	callZero := GoInt(zero)
	assert.True(t, callZero == 0)
	assert.Nil(t, err)

	Finalize()

	defer os.Remove("./hello.py")
	defer os.Remove("./hello.pyc")
}

func TestCompleteWorkFlowWithSubInterpreter(t *testing.T) {
	testFile, err := os.OpenFile("./hello.py", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	assert.Nil(t, err)
	testFile.Write([]byte("def minus(x,y):\n"))
	testFile.Write([]byte("	   return x - y\n"))
	testFile.Write([]byte("def add(x,y):\n"))
	testFile.Write([]byte("    return x + y\n"))
	testFile.Write([]byte("def return_self(self):\n"))
	testFile.Write([]byte("    return self\n"))
	testFile.Write([]byte("def zero_arg():\n"))
	testFile.Write([]byte("    return 0\n"))
	testFile.Close()

	err = Initialize()
	assert.Nil(t, err)

	subret, err := Py_NewInterpreter()
	assert.Nil(t, err)

	InsertPackagePath("./")

	ret, err := CallFunc("hello", "minus", 1, 1)
	callInt := GoInt(ret)
	assert.True(t, callInt == 0)
	assert.Nil(t, err)

	ret, err = CallFunc("hello", "add", 1, 1)
	callInt = GoInt(ret)
	assert.True(t, callInt == 2)
	assert.Nil(t, err)

	self1, err := CallFunc("hello", "return_self", 1)
	callSelf1 := GoInt(self1)
	assert.True(t, callSelf1 == 1)
	assert.Nil(t, err)

	self2, err := CallFunc("hello", "return_self", "self")
	callSelf2 := GoStr(self2)
	assert.True(t, callSelf2 == "self")
	assert.Nil(t, err)

	zero, err := CallFunc("hello", "zero_arg")
	callZero := GoInt(zero)
	assert.True(t, callZero == 0)
	assert.Nil(t, err)

	Py_EndInterpreter(subret)
	// Do not neccessary to invoke Finalize()
	//Finalize()

	defer os.Remove("./hello.py")
	defer os.Remove("./hello.pyc")
}

func TestPyRun_SimpleString(t *testing.T) {
	Initialize()
	os.Remove("/tmp/TestPyRun_SimpleString")

	interr := PyRun_SimpleString("import os")
	assert.Equal(t, 0, interr)
	interr = PyRun_SimpleString("os.makedirs('/tmp/TestPyRun_SimpleString')")
	assert.Equal(t, 0, interr)

	dir, err := os.Stat("/tmp/TestPyRun_SimpleString") //os.Stat获取文件信息
	assert.Nil(t, err)
	assert.True(t, dir.IsDir())

	Finalize()
}
