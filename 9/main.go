package main

import (
	"fmt"
	"math"
)

func main() {
	ch1 := make(chan uint8)
	ch2 := make(chan float64)

	numbers := []uint8{
		4, 13, 11, 6, 44, 13, 11, 67,
	}
	go func() {
		for _, v := range numbers {
			ch1 <- v
		}
		close(ch1)
	}()

	go convertNums(ch1, ch2)

	convertedNumbers := make([]float64, 0, len(numbers))

	for v := range ch2 {
		convertedNumbers = append(convertedNumbers, v)
	}
	fmt.Printf("origin: %+v\n", numbers)
	fmt.Printf("converted: %+v\n", convertedNumbers)
}

func convertNums(ch1 <-chan uint8, ch2 chan<- float64) {
	for v := range ch1 {
		ch2 <- math.Pow(float64(v), 3)
	}
	close(ch2)
}
