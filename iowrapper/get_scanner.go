package iowrapper

import (
	"bufio"
	"os"
)

func GetScanner(path string) (*bufio.Scanner, *os.File) {
	f := OpenReadFile(path)

	s := bufio.NewScanner(f)

	s.Buffer(make([]byte, 0), 2048*2048)

	return s, f
}
