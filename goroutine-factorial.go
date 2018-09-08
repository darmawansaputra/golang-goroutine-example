package main

import (
	"fmt"
)

func main() {
	var n int
	maxNumber := 2

	fmt.Print("Enter the number: ")
	fmt.Scanf("%d", &n)

	state := make(chan int, n)

	for i := n; i > 0; i-- {
		state <- i
	}

	var numbers [2]int

	for i := 0; i < n-1; i++ {
		for j := 0; j < maxNumber; j++ {
			numbers[j] = <-state
		}

		fmt.Printf("+ Creating Thread #%d (%d * %d)\n", i+1, numbers[0], numbers[1])

		go calculate(state, numbers, i+1)
	}

	result := <-state
	fmt.Printf("\n= Final result of %d! is %d", n, result)
}

func calculate(state chan<- int, numbers [2]int, th int) {
	fmt.Printf("- Result Thread #%d is %d\n", th, numbers[0]*numbers[1])
	state <- numbers[0] * numbers[1]
}
