package main

import (
	"fmt"
	"sync"
	"time"
)

func f(param string) {
	for i := 0; i < 5; i++ {
		fmt.Println(param, ":", i)
		time.Sleep(time.Millisecond)
	}
}

func main() {
	wg := sync.WaitGroup{}
	f("direct")

	go f("Hello")
	go f("World!")
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(msg string) {
			defer wg.Done()
			fmt.Println(msg)
		}(fmt.Sprintf("%s:%d", "going", i))
	}

	// time.Sleep(time.Second)
	wg.Wait()
	fmt.Println("done")
	time.Sleep(time.Millisecond * 1111)

}
