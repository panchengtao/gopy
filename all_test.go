package gopython

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestInsertExtraPackagePath(t *testing.T) {
	Initialize()
	InsertExtraPackagePath("./")
	Finalize()
}

// TODO:目前无法动态注入包搜索路径
func TestAll(t *testing.T) {
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

	//_, err = InsertExtraPackagePath("/home/panchengtao/go/src/pythoninvoker/python3.6")
	//assert.Nil(t, err)

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
