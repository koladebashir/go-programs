// An implementation of the word-count (wc) terminal utility
package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)
func main() {
	var nfile int = len(os.Args[1:])
	var lc_sum, wc_sum, cc_sum int
	
	if os.Args[1] == "-l" {

		for _, fname := range os.Args[2:] {
			file, err := os.Open(fname)
			var lc, wc, cc int
	
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
	
			scan := bufio.NewScanner(file)
	
			for scan.Scan() {
				wc += len(strings.Fields(scan.Text()))
				cc += len(scan.Text())
				lc++
			}
	
			if nfile > 2 {
				lc_sum += lc
				wc_sum += wc
				cc_sum += cc
			}
			file.Close()
			fmt.Printf("%8d %s\n", lc, fname)
		}
		if nfile > 2 {
			fmt.Printf("%8d total\n", lc_sum)
		}
	} else {
		for _, fname := range os.Args[1:] {
			file, err := os.Open(fname)
			var lc, wc, cc int
	
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
	
			scan := bufio.NewScanner(file)
	
			for scan.Scan() {
				wc += len(strings.Fields(scan.Text()))
				cc += len(scan.Text())
				lc++
			}
	
			if nfile > 1 {
				lc_sum += lc
				wc_sum += wc
				cc_sum += cc
			}
			file.Close()
			fmt.Printf("%8d %8d %8d %s\n", lc, wc, cc, fname)
		}
		if nfile > 1 {
			fmt.Printf("%8d %8d %8d total\n", lc_sum, wc_sum, cc_sum)
		}
	}
	
}