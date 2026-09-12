package main

import (
	"fmt"
	"time"
)

func main() {
	goroutinesCount := 10

	wg := NewWaitGroup()
	err := wg.Add(goroutinesCount)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i := range goroutinesCount {
		go func() {
			time.Sleep(time.Second * 2)
			fmt.Printf("Goroutine %d\n", i)
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Application is done")
}
