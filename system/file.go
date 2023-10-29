package system

import "os"

func FileExists(file string) bool {
	_, err := os.Stat(file)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}
