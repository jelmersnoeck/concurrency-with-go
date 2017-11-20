package main

import (
	"fmt"
	"time"
)

func main() {
	for _, v := range []int{1, 2, 3, 4, 5, 6} {
		go func() {
			fmt.Println(v)
		}()
	}

	time.Sleep(time.Second)
}
