// Write a program that defines a string variable called message
// with the value "Hi 👧🏽 and 👦🏽" and prints the fourth rune in it as a character, not a number.

package main

import "fmt"

func main() {
	var message string = "Hi 👧🏽 and 👦🏽"
	messageInRunes := []rune(message)
	fmt.Println(string(messageInRunes[3]))
}