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

const (
	// defaultLineLengthLimit defines how many bytes of a given line we'll take into consideration when comparing.
	// This protects us from OOM when reading an endless line.
	defaultLineLengthLimit = 100 * 1024 * 1024 // 100 MiB
)

// readFile opens the given file for reading and returns a reader and a closing function.
func readFile(path string, lineLimit int) (r *bufio.Reader, closeFn func() error, err error) {
	if path == "" {
		closeFn = func() error { return nil }
		r = bufio.NewReaderSize(os.Stdin, lineLimit)
		return r, closeFn, nil
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	r = bufio.NewReaderSize(file, lineLimit)
	return r, file.Close, nil
}

// outputUnique reads from the provided reader and outputs all unique lines.
func outputUnique(r *bufio.Reader, trim bool) error {
	// This map will hold the hashes of unique lines.
	m := make(map[uint64]struct{})

	var line []byte
	var partial bool
	var err error
	var hash uint64
	for {
		line, partial, err = r.ReadLine()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return fmt.Errorf("failed to read from file: %w", err)
		}
		if trim {
			line = bytes.TrimSpace(line)
		}
		hash = xxhash.Sum64(line)

		if _, exists := m[hash]; exists {
			// Not unique. Drain the rest of the line, if any, and continue.
			for partial {
				_, partial, err = r.ReadLine()
				if err == io.EOF {
					return nil
				}
				if err != nil {
					return fmt.Errorf("failed to read from file: %w", err)
				}
			}
			continue
		}

		// It's unique! We'll add it to the map, and we'll output it.
		m[hash] = struct{}{}
		fmt.Print(string(line))
		// The line was too long to read, so we only got the first `lineLengthLimit` bytes of it.
		// We still need to read through the rest.
		for partial {
			line, partial, err = r.ReadLine()
			if err == io.EOF {
				return nil
			}
			if err != nil {
				return fmt.Errorf("failed to read from file: %w", err)
			}
			// TODO There is a potential issue here with multibyte UTF-8 characters that can be split apart.
			fmt.Print(string(line))
		}
		fmt.Println()
	}
}

func main() {
	flag.Usage = func() {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "Usage of unique:\n\n")
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "When no arguments are given %s reads from the standard in.\n\n", os.Args[0])
		flag.PrintDefaults()
	}

	var filePath string
	var trim bool
	var lineLengthLimit int
	flag.StringVar(&filePath, "f", "", "path to the file to process")
	flag.BoolVar(&trim, "t", false, "trim whitespace from each line (default false)")
	flag.IntVar(&lineLengthLimit, "ll", defaultLineLengthLimit, "limit the length of each line being processed, ignoring any data beyond that length (values under 16 are ignored)")
	flag.Parse()

	reader, closeFn, err := readFile(filePath, lineLengthLimit)
	if err != nil {
		log.Fatalf("Failed to read from file '%s'. Error: %s\n", filePath, err.Error())
	}
	defer func() { _ = closeFn() }()

	err = outputUnique(reader, trim)
	if err != nil {
		log.Fatalf("Failed to process all data. Error: %s\n", err)
	}
}
