package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Rate pizza between 1 and 5 :")

	//comma ok syntax / err ok / try catch alternative
	input, _ := reader.ReadString('\n') //put _ at err to avoid handling error, similar for input put _
	fmt.Println("Thanks for rating, ", input)

	fmt.Printf("Type of this rating is %T \n", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating : ", numRating+1)
	}
}
