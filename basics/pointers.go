package main

import "fmt"

func testptr() {
	n := 5
	fmt.Println("Value of n ", n)
	fmt.Println("Address of n ", &n)
	var ptrN *int
	ptrN = &n
	fmt.Println("Value of ptrN ", ptrN)
	fmt.Println("Value at &of ptrN ", *ptrN)
}
