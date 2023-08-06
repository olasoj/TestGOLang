package main

import (
	"fmt"
	"math/rand"
	"time"

	//"sort"
	"runtime"
)

const size = 1_000_000_000

var data []float64

func main() {
	fmt.Println("Application started... ")
	start := time.Now()
	fillData()
	exeTime(start, "Data filling took: ")

	start = time.Now()
	isSliceSorted(data)
	exeTime(start, "Execution of solution took: ")

	start = time.Now()
	isSorted3(data)
	exeTime(start, "Execution of concurrent solution took: ")

}

func exeTime(start time.Time, msg string) time.Duration {
	end := time.Since(start)
	fmt.Println(msg, end)
	return end
}

func fillData() {
	data = make([]float64, size)
	for i := 0; i < size; i++ {
		data[i] = 100.0 * rand.Float64()
	}
}

func isSorted3(data []float64) bool {

	numberOfItems := len(data)
	ch := make(chan bool)
	numSegments := runtime.NumCPU()
	segmentSize := int(float64(numberOfItems) / float64(numSegments))

	// Launch numSegments goroutines
	for index := 0; index < numSegments; index++ {
		go isSegmentSorted(data, index*segmentSize,
			index*segmentSize+segmentSize, ch)
	}

	num := 0 // completed goroutines
	for {
		select {
		case value := <-ch: // Blocks until a goroutine puts a bool into the
			// channel
			if value == false {
				return false
			}
			num += 1
			if num == numSegments { // All goroutines have completed
				return true
			}
		}
	}
	return true
}

func isSegmentSorted(data []float64, i int, i2 int, ch chan bool) {
	for s := i + 1; i < i2; i++ {

		if data[s] > data[s+1] {
			fmt.Printf("%f is greater than %f\n", data[s], data[s+1])
			ch <- false
		}
	}

	ch <- true
}

func isSliceSorted[T int | float64](sliceInput []T) bool {
	length := len(sliceInput)

	if length < 2 {
		return true
	}

	var i int
	for i = 0; i < (length - 1); i++ {
		if sliceInput[i] > sliceInput[i+1] {
			return false
		}
	}

	return true
}
