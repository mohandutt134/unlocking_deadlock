// +build OMIT

package main

import (
	"fmt"
	"sync"
)

// START OMIT
var Wait sync.WaitGroup
var Counter int = 0

func main() {

	for routine := 1; routine <= 10; routine++ {
		go Routine(routine)
		Wait.Add(1)
	}

	Wait.Wait()
	fmt.Printf("Final Counter: %d\n", Counter)
}

func Routine(id int) {
	Counter += 1
	fmt.Println(Counter)
	Wait.Done()
}

//STOP OMIT
