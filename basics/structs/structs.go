package main

import "fmt"

// struct name User has 'U' suggesting that the struct is public an can be exported, same for struct attributes
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {

	fmt.Println("Structs in Go")

	bob := User{"bob", "bob@bob.com", false, 15}
	fmt.Println(bob)

	fmt.Printf("user : %+v\n", bob)

	fmt.Printf("%v email address : %v\n", bob.Name, bob.Email)

}
