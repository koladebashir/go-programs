// Start a new program. In main, declare an int variable called total. Write a for loop that uses a variable
// named i to iterate from 0 (inclusive) to 10 (exclusive). The body of the for loop should be as follows:
//     total := total + i
//     fmt.Println(total)
// After the for loop, print out the value of total. What is printed out? What is the likely bug in this code?

package main

import "fmt"

func main() {
	var total int
	for i := 0; i < 10; i++{
		total := total + i		// Bug: There is variable shadowing occuring on this line
		fmt.Println(total)		// Correct syntax: total = total + i or total += i
	}
	fmt.Println(total)
}

