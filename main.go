package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

// readFile opens the given file for reading and returns a reader and a closing
// function.
func readFile(fPath string) (r *bufio.Reader, closeFn func() error, err error) {
	file, err := os.Open(fPath)
	if err != nil {
		return nil, nil, err
	}
	return bufio.NewReader(file), file.Close, nil
}

func main() {
	flag.Usage = func() {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n\n", os.Args[0])
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "When no arguments are given %s reads from the standard in.\n", os.Args[0])
		flag.PrintDefaults()
	}

	var fPath string
	flag.StringVar(&fPath, "f", "", "path to the file to process")
	flag.Parse()

	var reader *bufio.Reader
	if fPath != "" {
		r, closeFn, err := readFile(fPath)
		if err != nil {
			log.Fatalf("Failed to read from file '%s'. Error: %s\n", err.Error())
		}
		reader = r
		defer closeFn()
	} else {
		reader = bufio.NewReader(os.Stdin)
	}

	// This map will hold the unique lines
	m := make(map[string]struct{})

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		// If it's not already in the map we'll output it and add it.
		if _, ok := m[string(line)]; !ok {
			m[string(line)] = struct{}{}
			fmt.Println(string(line))
		}
	}
}
