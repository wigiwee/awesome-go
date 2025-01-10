package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	fmt.Println("Go context")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ctx = enrichContext(ctx)

	go processRequest(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("Oh no request timeout, took more than 2 sec")
	}

	time.Sleep(2 * time.Second)

}

func processRequest(ctx context.Context) {

	rId := ctx.Value("request-id")
	fmt.Println(rId)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context Done")
			return
		default:
			fmt.Println("context not done yet, processing!!!")
		}
		time.Sleep(500 * time.Millisecond)
	}
}
func enrichContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "request-id", "12345")
}
