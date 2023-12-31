package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func pingGenerator(c chan string) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		c <- "ping"
		time.Sleep(time.Second * 1)
	}
}

func output(c chan string) {
	defer wg.Done()
	for {
		value := <-c
		fmt.Println(value)
	}
}

func main() {
	var c chan string
	c = make(chan string)
	wg.Add(2)
	go pingGenerator(c)
	go output(c)
	wg.Wait()
}
