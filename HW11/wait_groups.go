package main

import (
	"fmt"
	"sync"
)

func worker(k int) {
	fmt.Printf("Worker %d starting\n", k)

	fmt.Printf("Worker %d done\n", k)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func(n int) {
			defer wg.Done()
			worker(n)
		}(i)
	}

	wg.Wait()
}
