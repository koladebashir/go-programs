// Write two functions. The UpdateSlice function takes in a []string and a string. 
// It sets the last position in the passed-in slice to the passed-in string. 
// At the end of UpdateSlice, print the slice after making the change. 
// The GrowSlice function also takes in a []string and a string. 
// It appends the string onto the slice. 
// At the end of GrowSlice, print the slice after making the change. 
// Call these functions from main. Print out the slice before each function is called and after each function is called. 
// Do you understand why some changes are visible in main and why some changes are not?

package main

import "fmt"

func UpdateSlice(s []string, str string) {
	s[len(s)-1] = str
	fmt.Println(s)
}

func GrowSlice(s []string, str string) {
	s = append(s, str)
	fmt.Println(s)
}

func main() {
	sliceOfStrings := []string{"Hello", "World!"}
	fmt.Println(sliceOfStrings)
	UpdateSlice(sliceOfStrings, "Bye!")
	fmt.Println(sliceOfStrings)

	sliceOfStringsInJapanese := []string{"Konnichiwa", "Sekai!"}
	fmt.Println(sliceOfStringsInJapanese)
	GrowSlice(sliceOfStringsInJapanese, "Sayonera!")
	fmt.Println(sliceOfStringsInJapanese)
}