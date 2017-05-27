// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup

// START OMIT
func main() {

	c := make(chan int, 10)

	wg.Add(1)
	go func(c chan int) {
		for i := 0; i < 15; i++ {
			c <- int(rand.Int31n(100))
		}
		close(c)
		wg.Done()
	}(c)

	wg.Wait()
	for i := range c {
		fmt.Printf("%d\n", i)
	}
}

// STOP OMIT
