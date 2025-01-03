package main

import "fmt"

func main() {
	fmt.Println("Welcome to introduction to memory pointers")

	var one int = 2
	var ptr *int = &one

	fmt.Println("value of pointer is (memory address)", ptr)

	myNumber := 25

	var addr = &myNumber
	fmt.Println("value at pointer location is ", *addr)

	*addr = *addr + 2
	fmt.Println("value at pointer location after modification is ", *addr)

}
