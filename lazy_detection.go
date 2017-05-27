// +build OMIT

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// START OMIT
func main() {

	c := make(chan int)
	for i := 1; i <= 10; i++ {
		wg.Add(1)

		go func(c chan int, id int) {
			val := <-c
			fmt.Printf("%d-->", val)
			time.Sleep(time.Duration(100) * time.Millisecond)
			c <- id
		}(c, i)
	}
	c <- 0

	wg.Wait()
	fmt.Printf("%d-->%d", <-c, 0)
}

// STOP OMIT
