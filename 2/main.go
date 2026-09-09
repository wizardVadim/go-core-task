package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const sliceLength = 10
const maxNum = 100
const element = 11
const idx = 3

func main() {
	slice, err := randomIntSlice(sliceLength, maxNum)
	if err != nil {
		panic(err)
	}

	fmt.Printf("original slice: %v\n", slice)

	exampledSlice := sliceExample(slice)
	fmt.Printf("exampled slice: %v\n", exampledSlice)

	addedElSlice := addElements(slice, element)
	fmt.Printf("slice after adding an element: %v\n", addedElSlice)

	copiedSlice := copySlice(slice)
	fmt.Printf("original slice before: %v\n", slice)
	fmt.Printf("copied slice before: %v\n", copiedSlice)
	copiedSlice[idx] = maxNum
	fmt.Printf("original slice after: %v\n", slice)
	fmt.Printf("copied slice after: %v\n", copiedSlice)

	removedElSlice := removeElement(slice, idx)
	fmt.Printf("slice after removing an element: %v\n", removedElSlice)

	fmt.Printf("original slice after operations: %v\n", slice)
}

func randomIntSlice(length int, max int64) ([]int64, error) {
	slice := make([]int64, 0, length)
	for i := range length {
		_ = i
		val, err := rand.Int(rand.Reader, big.NewInt(max))
		if err != nil {
			return []int64{}, err
		}
		slice = append(slice, val.Int64())
	}
	return slice, nil
}

func sliceExample(in []int64) []int64 {
	if len(in) == 0 {
		return []int64{}
	}

	out := make([]int64, 0, len(in)/2)

	for _, v := range in {
		if v%2 == 0 {
			out = append(out, v)
		}
	}

	return out
}

func addElements(in []int64, value int64) []int64 {
	out := make([]int64, 0, len(in)+1)
	out = append(out, in...)
	out = append(out, value)
	return out
}

func copySlice(in []int64) []int64 {
	out := make([]int64, len(in))

	//copy(to, from) можно использовать
	for i, v := range in {
		out[i] = v
	}

	return out
}

func removeElement(in []int64, idx int) []int64 {
	// проверок на длину массива и тд делать не буду, так как в задании этого нет, должен вовращаться только слайс
	out := make([]int64, 0, len(in)-1)
	out = append(out, in[:idx]...)
	out = append(out, in[idx+1:]...)

	return out
}
