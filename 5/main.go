package main

import "fmt"

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	isAcross, res := isAcrossSlices(a, b)
	fmt.Printf("isAcross: %v, result slice: %+v\n", isAcross, res)
}

func isAcrossSlices(slice1 []int, slice2 []int) (bool, []int) {
	if len(slice1) == 0 || len(slice2) == 0 {
		return false, []int{}
	}
	index := make(map[int]struct{})

	for _, v := range slice2 {
		index[v] = struct{}{}
	}

	var length int
	if len(slice1) < len(slice2) {
		length = len(slice1)
	} else {
		length = len(slice2)
	}

	result := make([]int, 0, length)

	for _, v := range slice1 {
		_, exists := index[v]
		if exists {
			result = append(result, v)
		}
	}

	return len(result) > 0, result
}
