package fileUtils

import "os"

// IsExist returns a boolean indicating whether a file or directory exist.
func IsExist(filepath string) bool {
	_, err := os.Stat(filepath)
	if err == nil {
		return true
	}
	return !os.IsNotExist(err)
}

//IsDir returns a boolean indicating whether s file id a directory.
func IsDir(filepath string) bool {
	f, err := os.Stat(filepath)
	if err != nil {
		return false
	}
	return f.IsDir()
}
