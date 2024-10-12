// This is an implementation of the cat terminal utility using Go
package main

import (
	"os"
	"log"
	"fmt"
	"io"
)
// Helper function to open a file and return a closure to close the file
func getFile(name string) (*os.File, func(), error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, nil, err
	}

	return file, func() {
		file.Close()
	}, nil
}

func main() {
	file, closer, err := getFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer closer()

	data := make([]byte, 2048)
	for{
		n, err := file.Read(data)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			fmt.Println("")
			break
		}
		os.Stdout.Write(data[:n])
	}
}