package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Welcome to arrays in golang")

	var fruits [4]string

	fruits[0] = "apple"
	fruits[1] = "orange"
	fruits[2] = "peach"
	fruits[3] = "banana"

	fmt.Println("fruits list with ", len(fruits), " fruits: ", fruits)
	fmt.Printf("Type of fruit is %T", fruits)

	var vegetables = [3]string{"cabbage", "tomato", "mushroom"}
	fmt.Println(vegetables)

	fmt.Println("==============Slices==============")

	var fruitSlice = []string{"orange"}
	fmt.Printf("Type of fruitSlice is %T", fruitSlice)
	// fruitSlice[0] = "apple"	//not valid will throw exception
	fmt.Println(fruitSlice)

	fruitSlice = append(fruitSlice, "mango", "bananas")

	fmt.Println(fruitSlice)

	// fruitSlice = append(fruitSlice[1:])
	fruitSlice = append(fruitSlice[:1])
	fmt.Println(fruitSlice)

	highScores := make([]int, 4)

	highScores[0] = 22
	highScores[1] = 25
	highScores[2] = 29
	highScores[3] = 24

	fmt.Println("high scores before appending : ", highScores)

	// highScores[4] = 223 //this will throw an index out of bound excpetion
	highScores = append(highScores, 33, 23, 53, 22) //but this won't throw exception

	fmt.Println("high scores after appending : ", highScores)

	sort.Ints(highScores) //sorting

	fmt.Println("high scores after sorting : ", highScores)

	fmt.Println(sort.IntsAreSorted(highScores)) //return if slice is sorted

}
