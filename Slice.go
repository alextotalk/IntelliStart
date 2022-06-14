package main

import "fmt"

func sliceTest() {
	//slice := make([]int, 3, 3)
	slice := make([]int, 3, 3)
	slice[0] = 1
	slice[1] = 2
	slice[2] = 3
	fmt.Println(slice)
	Handle(slice)
	fmt.Println(slice)
	//var slice = []int{1, 2, 3, 4, 5}
	//fmt.Println(double(slice))

	//var list []int
	//print(list == nil, " - ;")
	//print(len(list) == 0, " - ;")
	//fmt.Println(append(list, 777))
	////
	////var list2 = []int{}
	////print(list2 == nil, " - ;")
	////print(len(list2) == 0, " - ;")
	//var slice = make([]int, 5, 10)
	//fmt.Println("Length=", len(slice), " - Capacity=", cap(slice))

}

func Handle(slice []int) {
	//slice[1] = 77
	//list := append(slice, 7777)
	//fmt.Println(list)
	//	It`s wrong we copy link so meaning parent slice changed.
	//	We need to crate new slice and copy meanings to it!
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	newSlice = append(newSlice, 777)
	fmt.Println(newSlice)
}

//func double(slice []int) []int {
//	//res:=[]int{} //wrong we recreate after write every time
//	//res := make([]int, 0, len(slice))
//	//for _, num := range slice {
//	//	res = append(res, num*2)
//	//}
//	//or
//	res := make([]int, len(slice))
//	for i, num := range slice {
//		res[i] = num * 2
//	}
//	return res
//}
