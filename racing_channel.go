package main

import (
	"fmt"
	"math/rand"
	"time"
)

func addCars(cars chan string, name string) {
	cars <- name
	fmt.Println("add to the channel", name)
}

func startRace(cars chan string) {
	fmt.Println("start race ........")
	for {
		car, open := <-cars
		fmt.Println(car)
		fmt.Println(open)
		if !open {
			break
		}

		fmt.Println(car, "car is running ....")
		time.Sleep(time.Duration(1+rand.Intn(5)) * time.Second)
	}
	fmt.Println("all cars have finished the race")
}

func RacingChannels() {
	cars := make(chan string, 2)

	go addCars(cars, "Frankie1")
	go addCars(cars, "Frankie2")

	time.Sleep(2 * time.Second)
	close(cars)
	time.Sleep(1 * time.Second)
	go startRace(cars)

	time.Sleep(8 * time.Second)
	fmt.Println("Race over!")

}
