package main

import "fmt"

func main() {

	//printing colorful text in go
	//specially useful for command line applications
	var Reset = "\033[0m"
	var Red = "\033[31m"
	var Green = "\033[32m"
	var Yellow = "\033[33m"
	var Blue = "\033[34m"
	var Magenta = "\033[35m"
	var Cyan = "\033[36m"
	var Gray = "\033[37m"
	var White = "\033[97m"

	println(Red + "This is Red" + Reset)
	println(Green + "This is Green" + Reset)
	println(Yellow + "This is Yellow" + Reset)
	println(Blue + "This is Blue" + Reset)
	println(Magenta + "This is Purple" + Reset)
	println(Cyan + "This is Cyan" + Reset)
	println(Gray + "This is Gray" + Reset)
	println(White + "This is White" + Reset)

	println("\033[33;1m This is Bright Yellow \033[0m")
	println("\033[33m This is Yellow \033[0m")

	mystring := "Hello World!"
	fmt.Println("Hello World")
	fmt.Println("\033[33m" + mystring)

	colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 34, "Hello")
	fmt.Println(colored)

}
