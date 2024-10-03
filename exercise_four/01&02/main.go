// Write a for loop that puts 100 random numbers between 0 and 100 into an int slice.
// Loop over the slice you created. For each value in the slice, apply the following rules:
// a. If the value is divisible by 2, print “Two!”
// b. If the value is divisible by 3, print “Three!”
// c. If the value is divisible by 2 and 3, print “Six!”. Don’t print anything else. d. Otherwise, print “Never mind”.
package main

import (
	"fmt"
	"math/rand"
)

func main(){
	randomNumbers := []int{}
	for i := 0; i < 100; i++ {
		randomNumbers = append(randomNumbers, rand.Intn(100))
	}

	for _, number := range randomNumbers {
		switch {
		case number%2 == 0 && number%3 == 0:
			fmt.Println("Six!")
		case number%2 == 0:
			fmt.Println("Two!")
		case number%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Never mind")
		}
	}
}