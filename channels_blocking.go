package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	ok := make(chan int)
	err := make(chan int)

	for _, v := range []int{1, 2, 3, 4, 5, 6} {
		go func(i int) {
			if i%2 == 0 {
				ok <- i // HL
			} else {
				err <- i // HL
			}
		}(v)
	}

	for i := 0; i < 6; i++ {
		select { // HL
		case o := <-ok: // HL
			fmt.Printf("OK: %d\n", o)
		case e := <-err: // HL
			fmt.Printf("ERR: %d\n", e)
		}
	}
}
