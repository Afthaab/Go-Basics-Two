package main

import (
	"sync"
)

type Drivers struct {
	driverCount int
	rw          sync.RWMutex
}

func main() {

}
