package io

import (
	"bufio"
	"os"
)

func GetScanner(path string) (*bufio.Scanner, *os.File) {
	f := OpenReadFile(path)

	s := bufio.NewScanner(f)

	s.Buffer(make([]byte, 0), 1024*1024)

	return s, f
}
