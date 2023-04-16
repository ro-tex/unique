// Unique is a command-line utility which ingests string values and outputs the unique ones.
// This is achieved by keeping track of the encountered values, which means that the consumed memory will grow with
// incoming unique values.

package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/cespare/xxhash"
)

// readFile opens the given file for reading and returns a reader and a closing function.
func readFile(fPath string) (r *bufio.Reader, closeFn func() error, err error) {
	file, err := os.Open(fPath)
	if err != nil {
		return nil, nil, err
	}
	return bufio.NewReader(file), file.Close, nil
}

// outputUnique reads from the provided reader and outputs all unique lines.
func outputUnique(r *bufio.Reader, tTrim bool) {
	// This map will hold the hashes of unique lines.
	m := make(map[uint64]struct{})

	var line []byte
	var err error
	var hash uint64
	for {
		line, _, err = r.ReadLine()
		if err == io.EOF {
			break
		}
		if tTrim {
			line = bytes.TrimSpace(line)
		}
		hash = xxhash.Sum64(line)
		// If it's not already in the map, we'll add it and output it.
		if _, ok := m[hash]; !ok {
			m[hash] = struct{}{}
			fmt.Println(string(line))
		}
	}
}

func main() {
	flag.Usage = func() {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n\n", os.Args[0])
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "When no arguments are given %s reads from the standard in.\n\n", os.Args[0])
		flag.PrintDefaults()
	}

	var fPath string
	var tTrim bool
	flag.StringVar(&fPath, "f", "", "path to the file to process")
	flag.BoolVar(&tTrim, "t", false, "trim whitespace from each line [DEFAULT: false]")
	flag.Parse()

	var reader *bufio.Reader
	if fPath != "" {
		r, closeFn, err := readFile(fPath)
		if err != nil {
			log.Fatalf("Failed to read from file '%s'. Error: %s\n", fPath, err.Error())
		}
		reader = r
		defer func() { _ = closeFn() }()
	} else {
		reader = bufio.NewReader(os.Stdin)
	}

	outputUnique(reader, tTrim)
}
