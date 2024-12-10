// Example from chapter 7 of Jon Bodner - Learning Go
package main

import "fmt"

type IntTree struct {
	value int
	left, right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{value: val}
	}

	if val < it.value {
		it.left = it.left.Insert(val)
	} else if val > it.value {
		it.right = it.right.Insert(val)
	}

	return it
}

func (it *IntTree) Contains(val int) bool {

	if it == nil {
		return false
	}

	if val == it.value {
		return true
	} else if val < it.value {
		if it.left == nil {
			return false
		}
		return it.left.Contains(val)
	} else {
		if it.right == nil {
			return false
		}
		return it.right.Contains(val)
	}
}

func main() {
	var it *IntTree

	it = it.Insert(10)
	it = it.Insert(2)
	it = it.Insert(12)
	fmt.Println("it.value:", it.value)
	fmt.Println("it.left.value:", it.left.value)
	fmt.Println("it.right.value:", it.right.value)

	fmt.Println(it.Contains(12))
}