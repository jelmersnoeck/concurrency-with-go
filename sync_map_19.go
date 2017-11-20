package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	monuments := map[string]string{
		"Brussels": "Atomium",
		"London":   "Big Ben",
		"New York": "Statue of Liberty",
		"Paris":    "Eiffel Tower",
	}
	reversed := &sync.Map{} // HL

	for k, v := range monuments {
		go func(k, v string) {
			reversed.Store(v, k) // HL
		}(k, v)
	}

	time.Sleep(time.Second)
	fmt.Println(reversed.Load("Atomium")) // HL
}
