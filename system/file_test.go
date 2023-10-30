package system

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/zddava/goext/assert"
)

func TestFileExists(t *testing.T) {
	assert.True(FileExists("file.go"), t)
	assert.False(FileExists("not_exists"), t)
}

func TestReadLine(t *testing.T) {
	f, err := os.Open("file.go")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := ReadLine(r)
		if IsReadError(err) {
			panic(err)
		}

		if IsEOF(err) {
			fmt.Println("end of file")
			break
		}

		fmt.Println(line)
	}

}
