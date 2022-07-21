package main

import (
	"fmt"
)

type Animal interface {
	getFeedPerMonth() int
	getProperties() (string, int)
}

func getDataFarm(farm Farm, period int) int {
	var totalFeedWeight int
	for _, v := range farm.Animals {
		name, weight := v.getProperties()
		feedPerMonth := v.getFeedPerMonth()
		//val := reflect.ValueOf(v)
		//name := val.FieldByName("Name").Interface().(string)
		//weight := val.FieldByName("Weight").Interface().(int)
		//fmt.Println(name, weight)
		fmt.Printf("Name - %v, weight - %vkg, need feed fo month - %vkg\n",
			name, weight, feedPerMonth)
		totalFeedWeight += feedPerMonth
	}
	fmt.Print(" -------------//------------- \n")
	fmt.Printf("Feed for one month - %vkg", totalFeedWeight)
	return totalFeedWeight * period
}

func main() {

	workingPeriod := 3
	donetskFarm := getFakeFarm()
	totalFeedWeight := getDataFarm(donetskFarm, workingPeriod)
	fmt.Print(" -------------//------------- \n")
	fmt.Printf("\nFeed for workig period(monthes) = %+v ", totalFeedWeight)

}
