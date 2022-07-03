package main

import "fmt"

func main() {
	arr := []int{4, 4, -2, -3, -2, -3, -5, 5}

	uniqueNumber := make(map[int]bool)
	result := make([]int, 0, len(arr))

	for _, value := range arr {

		if _, ok := uniqueNumber[value]; ok {
			continue
		}
		uniqueNumber[value] = false
		result = append(result, value)
	}

	fmt.Println(result)

}
