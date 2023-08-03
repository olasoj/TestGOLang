package main

import "fmt"

type Integer int

func (i Integer) String() string {
	return fmt.Sprintf("%d", i)
}
