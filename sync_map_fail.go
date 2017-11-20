package main

import (
	"fmt"
	"time"
)

func main() {
	monuments := map[string]string{
		"Brussels": "Atomium",
		"London":   "Big Ben",
		"New York": "Statue of Liberty",
		"Paris":    "Eiffel Tower",
	}
	reversed := map[string]string{}

	for k, v := range monuments {
		go func(k, v string) {
			for i := 0; i < 100; i++ {
				reversed[v] = k
			}
		}(k, v)
	}

	time.Sleep(time.Second)
	fmt.Println(reversed)
}
