package main

import "fmt"

func main() {
	input := []float64{17.3, 11.1, 9.9, 4.3, 12.6}
	res := MyFilter(input, func(num float64) bool {
		if num <= 10.0 {
			return true
		}
		return false
	})
	fmt.Println(res)
}

func MyFilter[E any](in []E, f func(E) bool) []E {

	var inputLen = len(in)
	var result = []E{}

	for i := 0; i < inputLen; i++ {
		var value = in[i]
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}
