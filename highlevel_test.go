package gopython

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestInsertPackagePath(t *testing.T) {
	testFile, err := os.OpenFile("./insert_package_path.py", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	assert.Nil(t, err)
	testFile.Write([]byte("def add(x,y):\n"))
	testFile.Write([]byte("    return x + y\n"))
	testFile.Close()

	Initialize()
	err = InsertPackagePath("./")
	assert.Nil(t, err)

	ret, err := CallFunc("insert_package_path", "add", 1, 1)
	callInt := GoInt(ret)
	assert.True(t, callInt == 2)
	assert.Nil(t, err)

	Finalize()

	defer os.Remove("insert_package_path.py")
	defer os.Remove("insert_package_path.pyc")
}

func TestCallFunc(t *testing.T) {
	Initialize()
	os.Remove("/tmp/TestCallFunc")
	_, err := CallFunc("os", "makedirs", "/tmp/TestCallFunc")
	assert.Nil(t, err)

	dir, err := os.Stat("/tmp/TestCallFunc")
	assert.Nil(t, err)
	assert.True(t, dir.IsDir())

	Finalize()
}
