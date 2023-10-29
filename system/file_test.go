package system

import (
	"testing"

	"github.com/zddava/goext/assert"
)

func TestFileExists(t *testing.T) {
	assert.True(FileExists("file.go"), t)
	assert.False(FileExists("not_exists"), t)
}
