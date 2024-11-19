package main

import "fmt"

func main() {
	p := Paypal{}
	s := Stripe{}
	b := BankTransfer{}

	ProcessPayment(p, 10.2)
	ProcessPayment(s, 10.2)
	ProcessPayment(b, 10.2)
	fmt.Println()
}