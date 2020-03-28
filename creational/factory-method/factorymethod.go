package factorymethod

import "fmt"

// Available payment methods
const (
	CreditCard = 1
	DebitCard = 2
	Paytm = 3
)

// PaymentMethod ... Base interface that would facilitate payments
type PaymentMethod interface {
	Pay(amount float32) string
}

// GetPaymentMethod ... Returns the suitable factory method
func GetPaymentMethod(method int) (PaymentMethod, error) {
	switch method {
	case CreditCard:
		return new(CreditCardPayment), nil
	case DebitCard:
		return new(DebitCardPayment), nil
	case Paytm:
		return new(PaytmPayment), nil
	default:
		return nil, fmt.Errorf("Payment method %d not supported", method)
	}
}

// CreditCardPayment ... Manage Credit card payments
type CreditCardPayment struct {}

// Pay ... Facilitate Credit card payment
func (c *CreditCardPayment) Pay(amt float32) (msg string) {
	msg = fmt.Sprintf("Your payment of %0.2f is sucessful", amt)
	return
}

// DebitCardPayment ... Manage Debit card payment
type DebitCardPayment struct {}

// Pay ... Facilitate Debit card payment
func (d *DebitCardPayment) Pay(amt float32) (msg string) {
	msg = fmt.Sprintf("Your payment of %0.2f is sucessful", amt)
	return
}

// PaytmPayment ... Manage Paytm payment
type PaytmPayment struct {}

// Pay ... Facilitate Paytm card payment
func (p *PaytmPayment) Pay(amt float32) (msg string) {
	msg = fmt.Sprintf("Your payment of %0.2f is sucessful", amt)
	return
}



