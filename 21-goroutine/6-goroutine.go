package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wgWorker := sync.WaitGroup{}
	wg.Add(1)
	//go1
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			wgWorker.Add(1)
			go func(id int) {
				defer wgWorker.Done()
				ch <- id
			}(i)
		}
		wgWorker.Wait()
		close(ch) //sending is finished over the channel ch

	}()

	for v := range ch {
		fmt.Println("recv", v)
	}
	wg.Wait()
}
