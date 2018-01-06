package iowrapper

import "os"

func OpenReadFile(path string) (*os.File) {
	f, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	return f
}
