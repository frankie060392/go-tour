package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func addCars(cars chan string, name string) {
	cars <- name
	fmt.Println("add to the channel", name)
}

func startRace(cars chan string) {
	fmt.Println("start race ........")
	// for {
	// 	car, open := <-cars
	// 	fmt.Println(car)
	// 	fmt.Println(open)
	// 	if !open {
	// 		break
	// 	}

	// 	fmt.Println(car, "car is running ....")
	// 	time.Sleep(time.Duration(1+rand.Intn(5)) * time.Second)
	// }

	for car := range cars {
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

func RunraceMilti() {

	lambo, ferari := make(chan string), make(chan string)

	var wg sync.WaitGroup
	wg.Add(3)
	go addCarsToChan(&wg, ferari, "ferari")
	// wg.Wait()
	// wg.Add(5)
	go addCarsToChan(&wg, lambo, "lambo")
	// wg.Wait()
	// wg.Add(10)
	go startRaceMulti(&wg, ferari, lambo)
	// Wait for the race to finish if do not use waitgroup
	// time.Sleep(25 * time.Second)
	wg.Wait()
	fmt.Println("Race over!")
}

func addCarsToChan(wg *sync.WaitGroup, c chan string, car string) {
	for i := 0; i < 5; i++ {
		c <- car
		fmt.Println(car, "add to channel")
	}
	wg.Done()
	close(c)
}

func startRaceMulti(wg *sync.WaitGroup, ferari chan string, lambo chan string) {
	for {
		select {
		case car, ok := <-ferari:
			if !ok {
				fmt.Println("Ferari hash closed")
				ferari = nil
				break
			}
			fmt.Println(car, "is running")
		case car, ok := <-lambo:
			if !ok {
				fmt.Println(car, "Lambo has closed")
				lambo = nil
				break
			}
			fmt.Println(car, "is running")
		}
		if ferari == nil && lambo == nil {
			break
		}
	}
	wg.Done()
	fmt.Println("All cars have finished the race!")
}
