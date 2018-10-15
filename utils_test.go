package gopython

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

// TODO:目前无法动态注入包搜索路径
func TestAll(t *testing.T) {
	testFile, err := os.OpenFile("./hello.py", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	assert.Nil(t, err)
	testFile.Write([]byte("def add(x,y):\n"))
	testFile.Write([]byte("    return x + y\n"))
	testFile.Write([]byte("def b(xixi,haha):\n"))
	testFile.Write([]byte("    return xixi + haha\n"))
	testFile.Close()

	err = Initialize()
	assert.Nil(t, err)

	//_, err = InsertExtraPackagePath("/home/panchengtao/go/src/pythoninvoker/python3.6")
	//assert.Nil(t, err)

	_, err = ImportModule("hello")
	assert.Nil(t, err)

	ret, err := CallFunc("hello", "b", "xixi", "haha")
	var callStr = GoStr(ret)
	assert.True(t, len(callStr) > 0)
	assert.Nil(t, err)

	ret, err = CallFunc("hello", "add", 1, 1)
	callInt := GoInt(ret)
	assert.True(t, callInt == 2)
	assert.Nil(t, err)

	Finalize()

	defer os.Remove("./hello.py")
	defer os.Remove("./hello.pyc")
}
