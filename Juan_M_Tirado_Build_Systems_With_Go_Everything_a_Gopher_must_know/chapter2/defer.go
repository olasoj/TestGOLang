package main

import "fmt"

func closeMsg() {
	fmt.Println("Closed!!!")
}

func main() {
	defer closeMsg()
	fmt.Println("Doing something...")
	defer fmt.Println("Certainly closed!!!")
	fmt.Println("Doing something else...")
}
