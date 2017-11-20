package main

import (
	"fmt"
	"time"
)

func main() {
	for _, v := range []int{1, 2, 3, 4, 5, 6} {
		go func(v int) {
			time.Sleep(time.Duration(v) * time.Second) // HL
			fmt.Println(v)
		}(v)
	}

	time.Sleep(3 * time.Second) // HL
}
