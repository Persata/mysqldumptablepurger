package iowrapper

import (
	"bufio"
	"os"
)

func GetWriter(path string) (*bufio.Writer, *os.File) {
	f := OpenWriteFile(path)

	w := bufio.NewWriter(f)

	return w, f
}
