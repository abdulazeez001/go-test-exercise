package main

import (
	"fmt"
	"math/rand"
	channel "section2/channel"
	sum "section2/goroutine"
	"time"
)

func main() {

	// ****** GOROUTINE TEST STARTS HERE ******** //
	size := 500
	numGoroutines := 8
	// Create a large array
	list := sum.CreateListForSize(size)
	sum := sum.SumArrayChunkConcurrently(list, numGoroutines)
	fmt.Printf("Sum: %d\n", sum)
	// ****** GOROUTINE TEST ENDS HERE ******** //

	// ****** CHANNEL TEST STARTS HERE ******** //
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	ch := make(chan int)
	randRange := 12
	go channel.Producer(ch, randRange)
	go channel.Consumer(ch)
	time.Sleep(time.Second * 10)
	close(ch)
	fmt.Println("Done!!!")
	// ****** CHANNEL TEST ENDS HERE ******** //
}
