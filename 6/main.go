package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go bitGenerator(ch)

	result := getRandomInt(ch, 100)
	fmt.Println(result)
}

func getRandomInt(bitCh chan int, max int) int {
	result := 0

	for result < max {
		bit := <-bitCh
		result = (result << 1) | bit
		if result >= max {
			result %= max
			break
		}
	}

	return result
}

func bitGenerator(bitCh chan int) {
	for {
		select {
		case bitCh <- 0:
		case bitCh <- 1:
		}
	}
}
