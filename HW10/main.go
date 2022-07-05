package main

import (
	"fmt"
)

type Animal interface {
	getFeedPerMonth() int
}

func getDataFarm(farm Farm, period int) int {
	var totalFeedWeight int
	for _, v := range farm.Animals {

		fmt.Println(v, "\n Name - %v ", v)
	}
	return totalFeedWeight
}
func main() {
	workingPeriod := 3
	donetskFarm := getFakeFarm()
	totalFeedWeight := getDataFarm(donetskFarm, workingPeriod)
	fmt.Printf("\nFeed for workig period = %+v ", totalFeedWeight)

}
