// +build OMIT

package main

import (
	"sync"
	"time"
)

// START OMIT
func main() {
	var l1, l2 sync.Mutex
	wg.Add(2)
	go func() {
		l1.Lock()
		time.Sleep(time.Duration(10) * time.Millisecond)
		l2.Lock()
		wg.Done()
	}()

	go func() {
		l2.Lock()
		time.Sleep(time.Duration(10) * time.Millisecond)
		l1.Lock()
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("Finished")
}

// STOP OMIT
