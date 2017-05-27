// +build OMIT

package main

import (
	"fmt"
	"sync"
	"time"
)

var l sync.RWMutex
var wg sync.WaitGroup
var Counter int = 0

// START OMIT
func main() {
	defer timeTrack(time.Now(), "Locking ")
	wg.Add(1)
	go func() {
		l.Lock()
		Counter = 5
		l.Unlock()
		wg.Done()
	}()

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func() {
			l.Lock()
			fmt.Printf("Reader %d get %d\n", i, Counter)
			time.Sleep(time.Duration(100) * time.Millisecond) // to simulate processing
			l.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
}

// STOP OMIT

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s", name, elapsed)
}
