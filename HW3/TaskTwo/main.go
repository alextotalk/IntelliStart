package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	input := "1 9 3 4 -5"
	var result string

	if len(input) < 2 {
		result = input
		fmt.Println(result)

		return
	}

	ArrOfStringElements := strings.Split(input, " ")

	convertToValidJson := fmt.Sprintf("[%v]", strings.Join(ArrOfStringElements, ","))

	maxLengthOfArrNumber := len(input)/2 + 1
	arrOfNumber := make([]int32, maxLengthOfArrNumber, maxLengthOfArrNumber)
	if err := json.Unmarshal([]byte(convertToValidJson), &arrOfNumber); err != nil {
		panic(fmt.Sprintf(
			"json is not valid: %v", err))
	}

	max := arrOfNumber[0]
	min := arrOfNumber[0]

	for _, value := range arrOfNumber {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}

	result = fmt.Sprintf("%v %v", max, min)

	fmt.Println(result)
}
