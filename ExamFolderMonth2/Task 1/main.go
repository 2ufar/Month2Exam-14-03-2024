package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomNumbers(n int, ch1 chan<- int) {
	rand.NewSource(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		ch1 <- rand.Intn(21) + 10
	}
	close(ch1)
}

func calculateSquares(ch1 <-chan int, ch2 chan<- int) {
	for num := range ch1 {
		ch2 <- num * num
	}
	close(ch2)
}

func printSquares(ch2 <-chan int) {
	for square := range ch2 {
		fmt.Printf("Squares are %d\n",square)
	}
}

func main() {
	rand.NewSource(time.Now().UnixNano())
	n := rand.Intn(6) + 5
	ch1 := make(chan int)
	ch2 := make(chan int)

	go generateRandomNumbers(n, ch1)
	go calculateSquares(ch1, ch2)
	go printSquares(ch2)

	<-time.After(2 * time.Second)
}