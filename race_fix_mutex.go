// +build OMIT

package main

import (
	"fmt"
	"sync"
)

var Wait sync.WaitGroup

// START OMIT
var l sync.Mutex
var Counter int = 0

func main() {

	for routine := 1; routine <= 10; routine++ {
		go Routine(routine)
		Wait.Add(1)
	}

	Wait.Wait()
	fmt.Printf("Final Counter: %d\n", Counter)
}

func Routine(id int) {
	l.Lock()
	value := Counter
	Counter = value + 1
	l.Unlock()
	Wait.Done()
}

// STOP OMIT
