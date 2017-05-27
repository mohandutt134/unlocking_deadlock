// +build OMIT

package main

import (
	"fmt"
	"sync"
	// START1 OMIT
	"sync/atomic"
)

// STOP1 OMIT

var Counter int32 = 0
var wg sync.WaitGroup

func main() {
	for routine := 1; routine <= 10; routine++ {
		wg.Add(1)
		go Routine(routine)
	}

	wg.Wait()
	fmt.Printf("Final Counter: %d\n", Counter)
}

// START2 OMIT
func Routine(id int) {
	atomic.AddInt32(&Counter, 1)
	wg.Done()
}

// STOP2 OMIT
