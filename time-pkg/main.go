package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Welcome to time package of Golang")

	presentTime := time.Now()

	presentHour := time.Now().Hour()

	fmt.Println("hour: ", presentHour)

	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006"))
	fmt.Println(presentTime.Format("01-02-2006 Monday"))
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.November, 28, 23, 44, 2, 223, time.UTC)

	fmt.Println(createdDate)

	fmt.Println(createdDate.Format("01-02-2006"))

}
