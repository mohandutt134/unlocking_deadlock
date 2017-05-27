// +build OMIT

package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	defer timeTrack(time.Now(), "dumb add")

	n := 1000000
	sum := 0

	Add(0, n, &sum)
	fmt.Println(sum)
}

func Add(start, end int, sum *int) {
	for i := start; i <= end; i++ {
		*sum += i
	}
}

// STOP OMIT

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s", name, elapsed)
}
