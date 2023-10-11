// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	c1 := "1"
// 	c2 := "2"
// 	c3 := "3"
// 	ch := make(chan string)
// 	wg := &sync.WaitGroup{}
// 	wg.Add(3)
// 	go func(c1 string, ch chan string) {

// 		wg.Done()
// 		ch <- c1
// 	}(c1, ch)
// 	go func(c2 string, ch chan string) {
// 		wg.Done()
// 		ch <- c2
// 	}(c2, ch)
// 	go func(c3 string, ch chan string) {
// 		wg.Done()
// 		ch <- c3
// 	}(c3, ch)

// 	for i := 0; i < 3; i++ {
// 		select { // evaluates only ones
// 		case afthab := <-ch:
// 			fmt.Println(afthab)
// 		case Messi := <-ch:
// 			fmt.Println(Messi)
// 		case Ronaldo := <-ch:
// 			fmt.Println(Ronaldo)
// 		}
// 		wg.Wait()
// 	}

// }
