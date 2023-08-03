package main

import (
	"fmt"
	"sync"
)

const number = 1000

var countValue int
var m sync.Mutex

func main() {
	var wg sync.WaitGroup
	wg.Add(number)

	for i := 0; i < number; i++ {
		go func() {
			defer wg.Done()
			m.Lock()
			defer m.Unlock()
			countValue++
		}()
	}

	wg.Wait()
	fmt.Printf("\ncountValue = %d", countValue)
}
