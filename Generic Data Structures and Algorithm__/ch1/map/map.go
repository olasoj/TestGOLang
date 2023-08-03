package main

import "fmt"

func MyMap[E, R any](slice []E, f func(E) R) []R {
	//Take the values of the slice

	var sliceSize = len(slice)
	newSlice := []R{}

	for i := 0; i < sliceSize; i++ {
		newSlice = append(newSlice, f(slice[i]))
	}

	return newSlice
}

func main() {
	slice := []int{1, 5, 2, 7, 4}
	result := MyMap(slice, func(i int) int { return i * i })

	fmt.Println(result)
}
