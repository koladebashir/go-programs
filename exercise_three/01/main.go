// Write a program that defines a variable named greetings of type slice of strings with 
// the following values: "Hello", "Hola", "नमस्कार", "こんにちは", and "Привіт". 
// Create a subslice containing the first two values; a second subslice with the second, third, and fourth values; 
// and a third subslice with the fourth and fifth values. Print out all four slices

package main

import "fmt"

func main() {
	var greeting []string = []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}

	first_subslice := greeting[:2]
	second_subslice := greeting[1:4]
	third_subslice := greeting[3:]

	fmt.Println("First subslice =", first_subslice)
	fmt.Println("Second subslice =", second_subslice)
	fmt.Println("Third subslice =", third_subslice)

}

