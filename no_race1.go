// +build OMIT

package main

import (
	"fmt"
	"sync"
	"time"
)

var Wait sync.WaitGroup
var l sync.Mutex
var Counter int = 0

func main() {

	for routine := 1; routine <= 2; routine++ {
		go Routine(routine)
		Wait.Add(1)
	}

	Wait.Wait()
	fmt.Printf("Final Counter: %d\n", Counter)
}

func Routine(id int) {
	l.Lock()
	value := Counter
	time.Sleep(time.Duration(id) * time.Second)
	for count := 0; count < 4; count++ {
		value++
	}
	Counter = value
	l.Unlock()

	Wait.Done()
}
