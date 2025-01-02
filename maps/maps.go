package main

import "fmt"

func main() {
	fmt.Println("maps in go")

	languages := make(map[string]string)

	languages[".js"] = "Javascript"
	languages[".java"] = "Java"
	languages[".py"] = "Python"
	languages[".rb"] = "Ruby"

	fmt.Println("List of all languages ", languages)

	fmt.Println("value for .js", languages[".js"])

	delete(languages, ".js")

	fmt.Println("List of all languages ", languages)

	//parsing maps
	fmt.Println("Parsing maps")
	for key, value := range languages {
		fmt.Printf("for key %v value is %v \n", key, value)
	}
}
