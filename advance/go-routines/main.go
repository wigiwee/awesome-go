package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	start := time.Now()

	username := fetchUser()
	// likes := fetchUserLikes(username)
	// match := fetchUserMatch(username)

	resChan := make(chan any, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	go fetchUserLikes(username, resChan, wg)
	go fetchUserMatch(username, resChan, wg)

	wg.Wait()
	close(resChan)
	//getting values from the channels
	for res := range resChan {
		fmt.Println("response", res)
	}
	// fmt.Println("likes:", likes, "match:", match)
	fmt.Println("took ", time.Since(start))

}

func fetchUser() string {
	time.Sleep(time.Millisecond * 100)
	return "BOB"
}

func fetchUserLikes(username string, resChan chan any, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 150)
	resChan <- 11 //writing values to channel
	//return 11			//old approach
}

func fetchUserMatch(username string, resChan chan any, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 100)
	resChan <- "ANNA"
	// return "ANNA"
}
