package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	const num = 1500
	var wg sync.WaitGroup
	counter := 0

	wg.Add(num)
	for i := 0; i < num; i++ {
		go func() {
			mu.Lock()
			counter++
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("count:", counter)
}
