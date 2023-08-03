package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var Musketeers = []string{
	"Athos", "Porthos", "Aramis", "D’Artagnan",
}

func GetMusketeer(id int) (string, error) {

	if id < 0 || id >= len(Musketeers) {
		return "", errors.New(fmt.Sprintf("Invalid id [%d]", id))
	}

	return Musketeers[id], nil
}

func main() {
	rand.Seed(time.Now().UnixNano())
	id := rand.Int() % 6
	musketeers, err := GetMusketeer(id)

	if err == nil {
		fmt.Printf("[%d] %s", id, musketeers)
	} else {
		fmt.Println(err)
	}
}
