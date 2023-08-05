package main

import (
	"fmt"
	"runtime"
	"sync"
)

const number = 1000

var countValue int
var m sync.Mutex

func main() {
	runtime.GOMAXPROCS(12)

	fmt.Printf("Number of cores %d", runtime.NumCPU())
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
	fmt.Printf("Number of cores %d", runtime.NumCPU())

}
