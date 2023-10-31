package main

import (
	"fmt"
	"time"
)

func main() {
	values := []int{1, 2, 3, 4, 5}

	for _, val := range values {

		go func(t int) {

			fmt.Printf("%d ", t)

		}(val)

	}

	time.Sleep(1000 * time.Millisecond)
}
