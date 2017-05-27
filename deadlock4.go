// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// START OMIT
var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go func() {
		time.Sleep(time.Duration(int(rand.Int31n(10))) * time.Millisecond)
		fmt.Println("Goroutine Finished")

		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Program Finished")
}

// STOP OMIT
