package main

import (
	"fmt"
	loggers "section1/interfaces"
	"section1/slices"
	"section1/swap"
)

func main() {

	// ****** SWAP TEST STARTS HERE ******** //
	x := 5
	y := 10
	fmt.Printf("Before swap: x = %d, y = %d\n", x, y)

	err := swap.Swap(&x, &y)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("After swap: x = %d, y = %d\n", x, y)
	}

	str1 := "hello"
	str2 := "world"
	fmt.Printf("Before swap: str1 = %s, str2 = %s\n", str1, str2)

	err = swap.Swap(&str1, &str2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("After swap: str1 = %s, str2 = %s\n", str1, str2)
	}

	// Test with nil pointers
	var a, b *int
	err = swap.Swap(a, b)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// ****** SWAP TEST ENDS HERE ******** //

	// ****** SLICE TEST ENDS HERE ******** //

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := slices.SumOfAllEvenNum(numbers)
	fmt.Printf("Sum of even numbers: %d\n", result)

	// Edge case: nil slice
	var nilSlice []int
	nilResult := slices.SumOfAllEvenNum(nilSlice)
	fmt.Printf("Sum of even numbers for nil slice: %d\n", nilResult)

	// Edge case: empty slice
	emptySlice := []int{}
	emptyResult := slices.SumOfAllEvenNum(emptySlice)
	fmt.Printf("Sum of even numbers for empty slice: %d\n", emptyResult)

	// ****** SLICE TEST ENDS HERE ******** //

	// ****** LOGGER TEST ENDS HERE ******** //
	// Create instances of FileLogger and ConsoleLogger
	fileLogger := &loggers.FileLogger{Destination: "log.txt"}
	consoleLogger := &loggers.ConsoleLogger{}

	// Create a slice of Logger interface to hold different loggers
	fileLogger.Log("File Logger\n")
	consoleLogger.Log("Console Logger")
	// ****** LOGER TEST ENDS HERE ******** //
}
