package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var lim int

const (
	WriteFour = "divisibleByFour"
	WriteTree = "divisibleByTree"
	WriteTwo  = "divisibleByTwo"
)

func randInt(numbers chan int) {
	for lim < 150 {
		number := rand.Intn(100)
		numbers <- number
	}
	close(numbers)
	fmt.Println(lim)
}

func Service(numbers <-chan int, done chan bool) {
	chanDivisibleTwo := make(chan int)
	chanDivisibleThree := make(chan int)
	chanDivisibleFour := make(chan int)
	go divisibleByTwo(chanDivisibleTwo, done)
	go divisibleByThree(chanDivisibleThree, done)
	go divisibleByFour(chanDivisibleFour, done)
	for {
		select {
		case num := <-numbers:
			if lim >= 150 {
				done <- true
				return
			}
			if num%2 == 0 {
				chanDivisibleTwo <- num
			}
			if num%3 == 0 {
				chanDivisibleThree <- num
			}
			if num%4 == 0 {
				chanDivisibleFour <- num
			}
		}
	}
}

func main() {
	numbers := make(chan int)
	done := make(chan bool)

	go randInt(numbers)
	go Service(numbers, done)

	<-done
}

func divisibleByTwo(number <-chan int, done chan bool) {
	var nums []int
	for {
		select {

		case num := <-number:
			if len(nums) > 50 {
				WriteFile(nums, WriteTwo)
				break
			}
			fmt.Println("Two ", num)
			nums = append(nums, num)
			lim++

		case <-time.After(10 * time.Microsecond):
			fmt.Println("timeout 1")
		}
	}
}

func divisibleByThree(number <-chan int, done chan bool) {
	var nums []int
	for {
		select {

		case num := <-number:
			if len(nums) > 50 {
				WriteFile(nums, WriteTree)
				break
			}
			fmt.Println("Tree ", num)
			nums = append(nums, num)
			lim++

		case <-time.After(30 * time.Millisecond):
			fmt.Println("timeout 2")
		}
	}
}

func divisibleByFour(number <-chan int, done chan bool) {
	var nums []int
	for {
		select {

		case num := <-number:
			if len(nums) > 50 {
				WriteFile(nums, WriteFour)
				break
			}
			fmt.Println("Four ", num)
			nums = append(nums, num)
			lim++
		case <-time.After(100 * time.Millisecond):
			fmt.Println("timeout 3")
		}
	}
}

func WriteFile(nums []int, c string) {

	f, err := os.Create("text.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for _, num := range nums {
		_, err = f.WriteString(fmt.Sprintf("%d %s\n", num, c))
		if err != nil {
			fmt.Printf("error writing string: %v", err)
		}
	}
}
