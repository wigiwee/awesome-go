package main

import "fmt"

// var jwtToken int := 3000 // := won't work here

const LoginToken string = "alskdjfl" // this variable is public because L of LoginToken is capital

func main() {
	var username string = "wigiwee"
	fmt.Println(username)
	fmt.Printf("variable is of type : %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type : %T \n", isLoggedIn)

	var smallValue uint8 = 255 // 256 would throw an exception since value is out of bound
	fmt.Println(smallValue)
	fmt.Printf("variable is of type : %T \n", smallValue)

	var smallfloat float32 = 258.6665488454
	fmt.Println(smallfloat)
	fmt.Printf("variable is of type : %T \n", smallfloat)

	var bigfloat float64 = 258.6665488454
	fmt.Println(bigfloat)
	fmt.Printf("variable is of type : %T \n", bigfloat)

	//default values and aliases
	var somevariable int
	fmt.Println(somevariable)
	fmt.Printf("variable is of type : %T \n", somevariable)

	//implicit type
	var website = "wigiwee.com"
	fmt.Println(website)
	// website = 3	// not allowed since website variable is treated as stirng

	//no var style
	numberOfUsers := 30000.0 // := does declaration and assignment but it only works inside a method
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("variable is of type : %T \n", LoginToken)

}
