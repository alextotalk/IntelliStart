package main

import "fmt"

//func one(n int) {
//	fmt.Println(n)
//	if n > 1 {
//		i := n - 1
//		one(i)
//	}
//
//}

func main() {
	fmt.Println("Hello world! Bro1!")
	slice := make([]int, 3, 3)
	slice[0] = 1
	slice[1] = 2
	slice[2] = 3
	fmt.Println(slice)
	handle(slice)
	fmt.Println(slice)
	sliceTest()
}
