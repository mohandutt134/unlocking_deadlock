// +build OMIT

package main

import (
	"fmt"
	"sync"
)

var Wait sync.WaitGroup

// START OMIT
func main() {
	c := make(chan int) // unbuffered channel
	for routine := 1; routine <= 10; routine++ {
		Wait.Add(1)
		go Routine(routine, c)
	}
	c <- 0 // Initiate data flow
	Wait.Wait()
	fmt.Printf("Final Counter: %d\n", <-c)
}

func Routine(id int, c chan int) {
	value := <-c
	counter := value + 1
	Wait.Done()
	c <- counter
}

// STOP OMIT
