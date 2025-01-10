package main

import (
	"context"
	"fmt"
	"time"
)

// func main() {
// 	start := time.Now()
// 	// ctx := context.Background()	//context of type Background
// 	ctx := context.WithValue(context.Background(), "foo", "bar") //context of type background with data stored alongside in key value manner
// 	userId := 34
// 	val, err := fetchUserData(ctx, userId)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("result:", val)
// 	fmt.Println("This took us :", time.Since(start))
// }

type Response struct {
	value int
	err   error
}

func fetchUserData(ctx context.Context, userId int) (int, error) {
	//200 milliseconds the maximum amount of time fetch function could return value
	foo := ctx.Value("foo") //accessing value in ctx
	fmt.Println(foo)
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel() //cancles context if fetch return value after 200 milliseconds
	resCh := make(chan Response)
	go func() {
		val, err := fetch()
		resCh <- Response{value: val, err: err}
	}()
	for {
		select {
		case <-ctx.Done():
			return -1, fmt.Errorf("fetch took more than 200 milliseconds")
		case resp := <-resCh:
			return resp.value, resp.err //the case where fetch() returned in time
		}
	}
}

// fetch some third party stuff which can be slow
func fetch() (int, error) {
	time.Sleep(time.Millisecond * 500)
	return 666, nil

}
