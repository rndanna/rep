package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup
var lim int

func randInt(numbers chan int) {
	for lim < 150 {
		number := rand.Intn(100)
		numbers <- number
	}
	close(numbers)
}

func Server(numbers chan int, done chan bool) {
	// rand.Seed(time.Now().UnixNano())

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	wg.Add(3)

	go Two(ch1)
	go Tree(ch2)
	go Four(ch3)

	for num := range numbers {
		// num := rand.Intn(100)

		if num%2 == 0 {
			select {
			case ch1 <- num:
			default:
			}
		}

		if num%3 == 0 {
			select {
			case ch2 <- num:
			default:
			}
		}

		if num%4 == 0 {
			select {
			case ch3 <- num:
			default:
			}
		}

	}

	wg.Wait()
	fmt.Println("done")
	done <- true
}

func Two(ch1 chan int) {
	var arr []int

	for {
		num := <-ch1
		arr = append(arr, num)

		if len(arr) >= 50 {
			fmt.Println("2 горутина:", arr)
			break
		}
	}

	wg.Done()
}

func Four(ch3 chan int) {
	var arr []int

	for {
		num := <-ch3
		arr = append(arr, num)

		if len(arr) >= 50 {
			fmt.Println("4 горутина:", arr)
			break
		}
	}

	wg.Done()
}

func Tree(ch2 chan int) {
	var arr []int

	for {
		num := <-ch2
		arr = append(arr, num)

		if len(arr) >= 50 {
			fmt.Println("3 горутина:", arr)
			break
		}
	}

	wg.Done()
}

func main() {
	numbers := make(chan int)
	done := make(chan bool)

	go randInt(numbers)
	go Server(numbers, done)

	<-done
}
