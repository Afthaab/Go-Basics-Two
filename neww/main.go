package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {

	// gives an empty container to put context values
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)

	defer cancel()

	ch := make(chan int)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(5 * time.Second)
		ch <- 10
	}()
	Recieve(ch, ctx)
	wg.Wait()
	// close(ch)

}

func Recieve(ch chan int, ctx context.Context) {
	select {
	case value := <-ch:
		fmt.Println(value)

	case cs := <-ctx.Done():
		fmt.Println("context", cs)
		return

	}
}
