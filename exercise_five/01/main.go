// create a simple calculator program having four operation functions
// add, sub, mul, div
// each function returns two variables, an int and an error
// for the div function, carry out a division by zero check and return an error errors.New("division by zero")

package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)
func add(x, y int) (int, error) { return x + y, nil}
func sub(x, y int) (int, error) { return x - y, nil}
func mul(x, y int) (int, error) { return x * y, nil}
func div(x, y int) (int, error) {
	var err error
	if y == 0 {
		err = errors.New("division by zero")
		return 0, err
	}
	return x / y, err
}

var operationMap = map[string]func(int, int) (int, error){
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func main() {
	testExpressions := [][]string{
		{"1", "+", "2"},
		{"201", "-", "1"},
		{"324", "*", "1"},
		{"2", "/", "1"},
		{"2", "1"},
		{"12", "/", "0"},
		{"2", "%", "1"},
	}

	for _, expression := range testExpressions {
		if len(expression) < 3 {
			fmt.Fprintln(os.Stderr, "Insufficient parameters to calculate")
			continue
		}
		x, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		op := expression[1]
		opFunc, ok := operationMap[op]
		if !ok {
			fmt.Printf("Operation %v not supported\n", op)
			continue
		}
		y, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		answer, err := opFunc(x, y)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		fmt.Println(expression, "=", answer)
	}
}