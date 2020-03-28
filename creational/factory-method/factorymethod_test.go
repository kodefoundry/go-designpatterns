package factorymethod

import "testing"

func TestGetPaymentMethodCreditCard(t *testing.T) {
	// <---------------- Arrange ---------------->
	pg, err := GetPaymentMethod(1)
	if pg == nil || err != nil {
		t.Fatal("A payment method for Credit card does not exists")
	}
	// <---------------- Act     ---------------->
	msg := pg.Pay(100.20)
	// <---------------- Assert  ---------------->
	if msg != "Your payment of 100.20 is sucessful" {
		t.Error("Credit card payment failed")
	}
	t.Log("LOG: ", msg)
}

func TestGetPaymentMethodDebitCard(t *testing.T) {
	// <---------------- Arrange ---------------->
	pg, err := GetPaymentMethod(2)
	if pg == nil || err != nil {
		t.Fatal("A payment method for Debit card does not exists")
	}
	// <---------------- Act     ---------------->
	msg := pg.Pay(110.00)
	// <---------------- Assert  ---------------->
	if msg != "Your payment of 110.00 is sucessful" {
		t.Error("Debit card payment failed")
	}
	t.Log("LOG: ", msg)
}

func TestGetPaymentMethodPaytm(t *testing.T) {
	// <---------------- Arrange ---------------->
	pg, err := GetPaymentMethod(3)
	if pg == nil || err != nil  {
		t.Fatal("A payment method for Paytm does not exists")
	}
	// <---------------- Act     ---------------->
	msg := pg.Pay(100.00)
	// <---------------- Assert  ---------------->
	if msg != "Your payment of 100.00 is sucessful" {
		t.Error("Debit card payment failed")
	}
	t.Log("LOG: ", msg)
}

func TestGetPaymentInvalid(t *testing.T) {
	// <---------------- Arrange ---------------->
	_, err := GetPaymentMethod(30)
	if err == nil {
		t.Error("A payment method for ID 30 does not exists")
	}
	t.Log("LOG: ", err)
}