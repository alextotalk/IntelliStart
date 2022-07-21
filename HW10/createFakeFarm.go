package main

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"math/rand"
	"time"
)

func getFakeFarm() Farm {

	animals := make([]Animal, 100, 100)
	sliceTypeAnimal := []Animal{Dog{}, Cat{}, Cow{}}
	rand.Seed(time.Now().Unix())

	for i, v := range animals {

		oneOfAnimals := sliceTypeAnimal[rand.Intn(3)]
		v = oneOfAnimals
		err := faker.FakeData(&v)
		if err != nil {
			fmt.Println(err)
		}
		animals[i] = v
	}
	return Farm{animals}
}
