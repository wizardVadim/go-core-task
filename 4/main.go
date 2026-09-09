package main

import "fmt"

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	result := originSliceElements(slice1, slice2)
	fmt.Println(result)
}

func originSliceElements(origin []string, other []string) []string {
	if len(origin) == 0 {
		return []string{}
	}

	index := make(map[string]struct{})

	for _, v := range other {
		index[v] = struct{}{}
	}

	result := make([]string, 0, len(origin))

	for _, v := range origin {
		_, exists := index[v]
		if !exists {
			result = append(result, v)
		}
	}

	return result
}
