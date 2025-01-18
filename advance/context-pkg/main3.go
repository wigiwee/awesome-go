package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// context
// 1. controlling timeout
// 2. cancelling go routine
// 3. passing metadata across go application

func main() {
	ctx := context.Background()
	exampleTimeout(ctx)
	exampleWithValue()
}

func exampleTimeout(ctx context.Context) {

	ctxWithTimeout, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	done := make(chan struct{})

	go func() {
		time.Sleep(3 * time.Second)
		close(done)
	}()
	select {
	case <-done:
		fmt.Println("Called the api")
	case <-ctxWithTimeout.Done():
		fmt.Println("oh no the timeout is expired ", ctxWithTimeout.Err())
	}
}

func exampleWithValue() {
	ctx := context.Background()

	ctxWithValue := context.WithValue(ctx, "userId", 344)

	userId := ctxWithValue.Value("userId")
	fmt.Printf("%t %v ", userId, userId)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("api response")
	case <-ctx.Done():
		fmt.Println("Oh no the context expired")
		http.Error(w, "request timeout", http.StatusRequestTimeout)
		return
	}
}
