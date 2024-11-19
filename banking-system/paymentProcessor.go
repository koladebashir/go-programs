// A short code on interfaces and how they are used
package main

import "fmt"

type PaymentProcessor interface {
	ProcessPayment(float64) string
}

type Paypal struct{}
func (p Paypal) ProcessPayment(amount float64) string {
	return fmt.Sprintf("Processed $%.2f with Paypal", amount)
}

type Stripe struct{}
func (s Stripe) ProcessPayment(amount float64) string {
	return fmt.Sprintf("Processed $%.2f with Stripe", amount)
}

type BankTransfer struct{}
func (b BankTransfer) ProcessPayment(amount float64) string {
	return fmt.Sprintf("Processed $%.2f with Bank transfer", amount)
}

func ProcessPayment(p PaymentProcessor, a float64) {
	fmt.Println(p.ProcessPayment(a))
}