package channel

import (
	"fmt"
	"math/rand"
	"time"
)

// Producer generates random numbers
//
//	and sends them to the channel
func Producer(channel chan<- int, randRange int) {
	for {
		num := rand.Intn(randRange)
		channel <- num
		time.Sleep(time.Millisecond * 500)
	}
}

// Consumer reads numbers from the channel,
// Calculates their square, and prints it
func Consumer(channel <-chan int) {
	for value := range channel {
		square := value * value
		fmt.Printf("%d * %d = %d\n", value, value, square)
	}
}
