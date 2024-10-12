// write a function called prefixer that has an input parameter of type string
// that returns a function with input parameter of type string that returns a string
// the returned function should prefix its input with the string passed into prefixer

package main

import "fmt"

func prefixer(prefix string) func(string) string {
	return func(s string) string {
		return prefix + " " + s
	}
}

func main() {
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob")) 	// should print Hello Bob
	fmt.Println(helloPrefix("Maria")) 	// should print Hello Maria
}