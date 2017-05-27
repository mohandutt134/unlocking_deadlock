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
	defer timeTrack(time.Now(), "start add")

	n := 1000000
	sum1 := 0
	sum2 := 0
	sum3 := 0
	wg.Add(3)

	go Add(0, n/3, &sum1)
	go Add(n/3+1, 2*n/3, &sum2)
	go Add(2*n/3+1, n, &sum3)

	wg.Wait()
	fmt.Println(sum1 + sum2 + sum3)
}

func Add(start, end int, sum *int) {
	for i := start; i <= end; i++ {
		*sum += i
	}
	wg.Done()
}

// STOP OMIT

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s", name, elapsed)
}
