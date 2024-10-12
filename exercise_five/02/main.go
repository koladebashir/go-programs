// Write a function called fileLen that has an input parameter of type string and returns an int and an error.
// The function takes in a filename and returns the number of bytes in the file.
// If there is an error reading the file, return the error.
// Use defer to make sure the file is closed properly.

package main

import (
	"os"
	"io"
	"fmt"
)

func fileLen(fileName string) (int, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	data := make([]byte, 2048)
	numberOfBytes := 0
	for {
		n, err := f.Read(data)
		if err != nil {
			if err != io.EOF {
				return 0, err
			}
			break
		}
		numberOfBytes += n

	}

	return numberOfBytes, nil
}

func main() {
	if len(os.Args) < 2 {
		return
	}
	n, err := fileLen(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Size: %v bytes\n", n)
}