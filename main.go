package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var fName string
	flag.StringVar(&fName, "f", "", "name of the file to process")
	flag.Parse()

	if fName == "" {
		flag.Usage()
		os.Exit(1)
	}

	file, err := os.Open(fName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// This map will hold the unique lines
	m := make(map[string]struct{})

	reader := bufio.NewReader(file)
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
