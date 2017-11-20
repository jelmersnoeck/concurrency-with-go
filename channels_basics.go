package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 1 // HL
	}()

	fmt.Println(<-c) // HL
}
