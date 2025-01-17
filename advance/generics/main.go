package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// normal method
// func add(a, b int) int {
// 	return a + b
// }

// generic Add functions but only supports int
// func add[T int](a, b T) T {
// 	return a + b
// }

// to support multiple types
// now works for all the defined types but there are a lot of types
// func add[T int | float32 | float64 | string | int8](a, b T) T {
// 	return a + b
// }

//to work around above problem

// create new type containing all other types
// ~ implies that it will accept any type which is int8 underneeth ie. integer8
type allTypes interface {
	~int8 | int16 | int | uint | uint16 | float32 | float64 | string
}

type integer8 interface {
	int8
}

// and we can use allTypes here
// func add[T allTypes](a, b T) T {
// 	return a + b
// }

// to avoid using allTypes using constraints
// constraints.Ordered contains all types
func add[T constraints.Ordered](a, b T) T {
	return a + b
}

func main() {

	fmt.Println(add(3, 5))
	fmt.Println(add(392.2, 32.234))
	fmt.Println(add("prefix_", "suffix"))
	fmt.Println(add(8, 4))

}
