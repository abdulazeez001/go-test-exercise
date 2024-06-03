package sum

import (
	"fmt"
	"sync"
)

func CreateListForSize(size int) []int {
	list := make([]int, size)
	for index := 0; index < size; index++ {
		list[index] = index + 1
	}
	fmt.Println(list)
	return list
}

func sumArrayChunk(list []int, start, end int, waitGroup *sync.WaitGroup, ch chan<- int) {
	defer waitGroup.Done()
	sum := 0
	for index := start; index < end; index++ {
		sum += list[index]
	}
	ch <- sum
}

func SumArrayChunkConcurrently(list []int, numGoroutines int) int {

	channel := make(chan int, numGoroutines)
	var waitGroup sync.WaitGroup

	chunkSize := (len(list) + numGoroutines - 1) / numGoroutines

	for index := 0; index < numGoroutines; index++ {
		start := index * chunkSize
		end := start + chunkSize
		if end > len(list) {
			end = len(list)
		}
		waitGroup.Add(1)
		go sumArrayChunk(list, start, end, &waitGroup, channel)
	}

	// Close the channel once all goroutines have finished
	go func() {
		waitGroup.Wait()
		close(channel)
	}()

	totalSum := 0
	for sum := range channel {
		totalSum += sum
	}

	return totalSum

}
