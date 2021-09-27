package util

import (
	"bufio"
	"os"
	"io"
	"strings"
)

// Read numbers and text from file

type In struct {
	scanner *bufio.Scanner
	file *os.File
}

// Factory method
func NewIn(path string) *In {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	return &In{file: f}
}

func ReadAllStrings(path string) ([]string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	str := string(data)

	return strings.Fields(str)
}