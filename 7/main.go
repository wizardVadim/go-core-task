package main

import (
	"fmt"
	"sync"
)

func main() {
	chSliceSize := 10
	chIterations := 100

	chSlice := make([]<-chan int, 0, chSliceSize)

	for range chSliceSize {
		currentChannel := make(chan int)
		chSlice = append(chSlice, currentChannel)
		go func() {
			for i := range chIterations {
				currentChannel <- i
			}
			close(currentChannel)
		}()
	}

	result := channelFromChannels(chSlice)
	slice := make([]int, 0)
	for value := range result {
		slice = append(slice, value)
	}
	fmt.Println(slice)
}

func channelFromChannels(channels []<-chan int) <-chan int {
	result := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(channels))
	for _, ch := range channels {
		go func() {
			defer wg.Done()
			for value := range ch {
				result <- value
			}
		}()
	}

	go func() {
		wg.Wait()
		close(result)
	}()
	return result
}
