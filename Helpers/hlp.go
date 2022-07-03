package hlp

import "fmt"

func MaxMin(a, b int) (max, min int) {
	if a > b {
		return a, b
	} else {
		return b, a
	}
}

func MaxMimInArr(arr []int) (max, min int) {
	max = arr[0]
	min = arr[0]

	for _, value := range arr {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return max, min
}
func PrettyDivisor() {
	fmt.Print(" -------------//------------- \n")
}
