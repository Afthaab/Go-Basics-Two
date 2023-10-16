package main

import (
	"fmt"
	"sync"
)

var myMap = make(map[int]int)
var wg sync.WaitGroup
var mutex = sync.Mutex{}

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			mutex.Lock()
			myMap[n] = n * n
			mutex.Unlock()
		}(i)
	}
	wg.Wait()

	fmt.Println(myMap)
}
