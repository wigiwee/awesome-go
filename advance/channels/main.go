package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("Channels in golang")

	myChan := make(chan int) //non buffered channel
	// myChan := make(chan int, 2) //buffered channel

	//will throw exception
	//because we can only push value to channel if someone is listening
	//and we can only listen on a channel if someone is pushing value to it
	//its a chicken egg situation
	// myChan <- 5           //pushing value to chan
	// fmt.Println(<-myChan) //getting value out of channel

	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println(<-myChan) //reading value from the channel
		fmt.Println(<-myChan) //reading value from the channel
		val, isChannelOpen := <-myChan
		fmt.Println(isChannelOpen, val)
	}(myChan, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		myChan <- 5 //adding value to channed
		myChan <- 6 //adding value to channed
		close(myChan)
	}(myChan, wg)

	//write only Channel (can close channel)
	go func(ch chan<- int, wg *sync.WaitGroup) {
		defer wg.Done()
		myChan <- 5 //adding value to channed
		myChan <- 6 //adding value to channed
		close(myChan)
	}(myChan, wg)

	//read only Channel (cannot close channel)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		myChan <- 5 //adding value to channed
		myChan <- 6 //adding value to channed
		close(myChan)
	}(myChan, wg)
	wg.Wait()
}
