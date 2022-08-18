package testutil

import (
	"bufio"
	"compress/gzip"
	"errors"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// Reads in data of various types from standard input, files and URLs.

var ErrEmpty = errors.New("argument is empty")

type In struct {
	reader     io.Reader
	scanner    *bufio.Scanner
	hasScanned bool
	hasNext    bool
}

// Factory method
// default read in lines
func NewInWithError(uri string) (*In, error) {
	r, err := newReader(uri)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(r)

	return &In{reader: r, scanner: scanner}, nil
}

func NewIn(uri string) *In {
	return NewInReadLines(uri)
}

func NewInReadWords(uri string) *In {
	r, err := newReader(uri)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(r)

	scanner.Split(bufio.ScanWords)

	return &In{reader: r, scanner: scanner}
}

func NewInReadLines(uri string) *In {
	r, err := newReader(uri)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(r)

	return &In{reader: r, scanner: scanner}
}

func newReader(uri string) (io.Reader, error) {
	if uri == "" {
		return nil, ErrEmpty
	}

	// first try to read file from local file system
	f, err := os.Open(uri)
	if err == nil {
		if strings.HasSuffix(uri, ".gz") {
			gz, err := gzip.NewReader(f)
			if err != nil {
				panic(err)
			}
			return gz, nil
		} else {
			return f, nil
		}
	} else {
		// URL from web
		resp, err := http.Get(uri)
		if err != nil {
			return nil, err
		}
		return resp.Body, nil
	}

}

func (in *In) ReadString() string {
	in.next()
	return in.scanner.Text()
}

func (in *In) ReadLine() string {
	in.next()
	return in.scanner.Text()
}

func (in *In) ReadInt() int {
	i, _ := strconv.Atoi(in.ReadString())
	return i
}

func (in *In) ReadFloat() float64 {
	f, _ := strconv.ParseFloat(in.ReadString(), 64)
	return f
}

func (in *In) IsEmpty() bool {
	return !in.HasNext()
}

func (in *In) HasNext() bool {
	if !in.hasScanned {
		in.hasNext = in.scanner.Scan()
		in.hasScanned = true
	}
	return in.hasNext
}

func (in *In) next() bool {
	if in.hasScanned {
		in.hasScanned = false
		return in.hasNext
	}
	return in.scanner.Scan()
}

func (in *In) ReadAll() string {
	data, err := io.ReadAll(in.reader)
	if err != nil {
		panic(err)
	}

	return string(data)
}

func (in *In) ReadAllLines() []string {
	ss := make([]string, 0)

	for in.HasNext() {
		ss = append(ss, in.ReadString())
	}

	return ss
}

func (in *In) ReadAllStrings() []string {
	str := in.ReadAll()

	return strings.Fields(str)
}

func (in *In) ReadAllInts() []int {
	strSlice := in.ReadAllStrings()
	n := len(strSlice)
	intSlice := make([]int, n)

	for i := 0; i < n; i++ {
		intSlice[i], _ = strconv.Atoi(strSlice[i])
	}

	return intSlice
}
