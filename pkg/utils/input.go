package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
)

func usage() {
	fmt.Printf("usage: %s <input>\n", os.Args[0])
	os.Exit(1)
}

func ReadInput() <-chan string {
	if len(os.Args) < 2 {
		usage()
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	out := make(chan string)

	go func() {
		defer file.Close()

		bufReader := bufio.NewReader(file)
		for {
			line, _, err := bufReader.ReadLine()
			if err == io.EOF {
				close(out)
				return
			}

			out <- string(line)
		}
	}()

	return out
}

func ReadInputByte() <-chan []byte {
	if len(os.Args) < 2 {
		usage()
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	out := make(chan []byte)

	go func() {
		defer file.Close()

		bufReader := bufio.NewReader(file)
		for {
			line, _, err := bufReader.ReadLine()
			if err == io.EOF {
				close(out)
				return
			}

			out <- slices.Clone(line)
		}
	}()

	return out
}
