package system

import (
	"bufio"
	"io"
	"os"
)

func FileExists(file string) bool {
	_, err := os.Stat(file)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func ReadLine(r *bufio.Reader) (string, error) {
	line, isprefix, err := r.ReadLine()
	for isprefix && err == nil {
		var next []byte
		next, isprefix, err = r.ReadLine()
		line = append(line, next...)
	}

	return string(line), err
}

func IsReadError(err error) bool {
	return err != nil && err != io.EOF
}

func IsEOF(err error) bool {
	return err != nil && err == io.EOF
}
