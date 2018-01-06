package iowrapper

import "os"

func OpenWriteFile(path string) (*os.File) {
	f, err := os.Create(path)

	if err != nil {
		panic(err)
	}

	return f
}
