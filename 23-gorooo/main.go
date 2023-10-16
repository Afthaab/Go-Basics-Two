package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)
	done := make(chan struct{})
	wg := sync.WaitGroup{}

	wgWorker := sync.WaitGroup{}

	wgWorker.Add(3)
	go func() {
		defer wgWorker.Done()
		time.Sleep(4 * time.Second)
		c1 <- "1"
		c1 <- "4"
	}()
	go func() {
		defer wgWorker.Done()
		time.Sleep(2 * time.Second)
		c2 <- "2"
	}()
	go func() {
		defer wgWorker.Done()
		c3 <- "3"
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		wgWorker.Wait()
		close(done)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {

			select {
			case x := <-c1:
				fmt.Println(x)
			case x := <-c2:
				fmt.Println(x)
			case x := <-c3:
				fmt.Println(x)
			case <-done:
				fmt.Println("work is finished")
				return
			}
		}
	}()
	wg.Wait()
}
