package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup //usually these are pointers
var mut sync.Mutex    //usually we create pointer of this
var signals = []string{"test"}

func main() {
	// greeter("Hello")
	// greeter("World")

	// fmt.Println()
	// fmt.Println()

	// //using go routing
	// go greeter("Hello")
	// greeter("world")
	weblist := []string{"https://google.com", "https://go.dev", "https://youtube.com", "https://fb.com"}

	//this approach takes long time
	// for _, web := range weblist {
	// 	getStatusCode(web)
	// }

	//implementing go routing in above example

	//this will return nothing since the method will be exited
	//before the thread could return

	for _, web := range weblist {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait() //holds the main method until wg.Done is fired in getstatus method

	//in the getStatusCode method multiple threads/goroutines might try to write to a single memory adderess at a same time, this might cause problems
	//mutex is used to solve this issue
	fmt.Println(signals)
	//waitGroup is used for solving above problem

}

func greeter(s string) {

	for range 5 {
		// time.Sleep(2 * time.Second)
		fmt.Println(s)

	}
}

func getStatusCode(endpoint string) int {
	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	}
	//locking resource to avoid concurrent write
	mut.Lock()
	signals = append(signals, endpoint)
	mut.Unlock()
	fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	return res.StatusCode
}
