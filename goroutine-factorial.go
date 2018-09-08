package main

import (
	"fmt"
)

func main() {
	// Define n for factorial
	var n int

	// Ask input number factorial to user
	fmt.Print("Enter the number: ")
	fmt.Scanf("%d", &n)

	// Create channel with n total buffer
	state := make(chan int, n)

	// Input number to state channel
	for i := n; i > 0; i-- {
		state <- i
	}

	// Int variable to help process
	var numbers [2]int

	// Looping with n - 1 thread will be created
	for i := 0; i < n-1; i++ {
		// Fill numbers array with each integer value from state channel
		for j := 0; j < 2; j++ {
			numbers[j] = <-state
		}

		// Creating thread
		fmt.Printf("+ Creating Thread #%d (%d * %d)\n", i+1, numbers[0], numbers[1])
		go calculate(state, numbers, i+1)
	}

	// Receive last value from state channel and print the result
	result := <-state
	fmt.Printf("\n= Final result of %d! is %d", n, result)
}

// Thread function to calculate number
func calculate(state chan<- int, numbers [2]int, th int) {
	fmt.Printf("- Result Thread #%d is %d\n", th, numbers[0]*numbers[1])
	state <- numbers[0] * numbers[1]
}
