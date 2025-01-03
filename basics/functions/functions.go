package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {

	//functions
	fmt.Println("Functions in golang")
	greeter()

	fmt.Println(adder(3, 2, 5, 2))

	//methods
	// method -> works with a struct/ obj

	fmt.Println("Methods in golang")

	bob := User{"Bob", "bob@gmail.com", true, 33}
	bob.GetStatus()
	bob.NewEmail()
	fmt.Println("user email is ", bob.Email) //above method changed email but it is not updated here

	//defer statement

	//multiple defer statements run in LIFO sequence
	defer fmt.Println("World defer1")
	defer fmt.Println("World defer2")
	fmt.Println("hello")
	fmt.Println("Before deferFunc() call")
	deferFunc()
	fmt.Println("After deferFunc() call")
}
func deferFunc() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func adder(values ...int) int {
	total := 0
	// values is of type slice
	fmt.Printf("%T \n", values)
	for _, value := range values {
		total += value
	}
	return total
}

func greeter() {
	fmt.Println("Hello")
}

// method -> works with a struct/ obj
func (u User) GetStatus() {
	fmt.Println("Is user active : ", u.Status)
}

func (u User) NewEmail() {
	u.Email = "test@gmail.com"
	fmt.Println("Email of this user is ", u.Email)
}
