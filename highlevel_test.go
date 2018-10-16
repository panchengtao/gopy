package gopython

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

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
