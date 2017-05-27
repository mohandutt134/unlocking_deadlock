// +build OMIT

package main

import (
	"fmt"
	"sync"
	"time"
)

// START OMIT
var l sync.Mutex
var Counter int = 0

func main() {
	for routine := 1; routine <= 10; routine++ {
		go Routine(routine)
	}
	time.Sleep(time.Duration(1) * time.Second)
	fmt.Printf("Final Counter: %d", Counter)
}

func Routine(id int) {
	// l.Lock()
	//fmt.Println(Counter)
	value := Counter
	time.Sleep(time.Duration(1) * time.Millisecond)
	Counter = value + 1
	fmt.Println(Counter)
	// l.Unlock()
}

// STOP OMIT
