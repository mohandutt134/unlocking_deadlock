// +build OMIT

package main

import (
	"sync"
)

// START OMIT
func main() {
	var wg sync.WaitGroup
	var l sync.Mutex
	l.Lock()
	wg.Add(1)
	go func() {
		l.Lock()
		wg.Done()
	}()
	wg.Wait()
}

// STOP OMIT
