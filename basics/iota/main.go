package main

import (
	"fmt"
)

// self incrementing value startin from 0
const (
	Monday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

const (
	A = iota + 334
	B
	C
)

const (
	Readable = 1 << iota // 1 << 0 = 001
	Writable
	Executable
)

func main() {

	fmt.Println("Week days")
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)
	fmt.Println(Thursday)
	fmt.Println(Friday)
	fmt.Println(Saturday)
	fmt.Println(Sunday)

	fmt.Println()

	fmt.Println("some values starting from 334")
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)

	fmt.Println()

	fmt.Println("flags values")
	fmt.Printf("%03b \n", Readable)
	fmt.Printf("%03b \n", Writable)
	fmt.Printf("%03b \n", Executable)

}
