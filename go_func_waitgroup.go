package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup // HL

	for _, v := range []int{1, 2, 3, 4, 5, 6} {
		wg.Add(1) // HL

		go func(v int) {
			defer wg.Done() // HL

			time.Sleep(time.Duration(v) * time.Second)
			fmt.Println(v)
		}(v)
	}

	wg.Wait() // HL
}
