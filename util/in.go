package util

import (
	"bufio"
	"os"
	"io"
	"strings"
	"compress/gzip"
)

// Reads in data of various types from standard input, files and URLs.

type In struct {
	file *os.File
	scanner *bufio.Scanner
}

// Factory method
// default read in lines
func NewIn(path string) *In {
	return NewInReadLines(path)
}

func NewInReadWords(path string) *In {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	var scanner *bufio.Scanner

	if strings.HasSuffix(path, ".gz") {
		gz, err := gzip.NewReader(f)
		if err != nil {
			panic(err)
		}
		
		scanner = bufio.NewScanner(gz)
	} else {
		scanner = bufio.NewScanner(f)
	}
	
	
	scanner.Split(bufio.ScanWords)

	return &In{f, scanner}
}

func NewInReadLines(path string) *In {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	
	scanner := bufio.NewScanner(f)

	return &In{f, scanner}
}

func (in *In) ReadString() string {
	return in.scanner.Text()
}

func (in *In) IsEmpty() bool {
	return !in.scanner.Scan()
}

func (in *In) ReadAllStrings() ([]string) {
	str := in.readAll()

	return strings.Fields(str)
}

func (in *In) readAll() string {
	data, err := io.ReadAll(in.file)
	if err != nil {
		panic(err)
	}

	return string(data)
}