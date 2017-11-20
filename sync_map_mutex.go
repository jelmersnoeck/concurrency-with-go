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
	reversed := map[string]string{}
	m := &sync.Mutex{} // HL

	for k, v := range monuments {
		go func(k, v string) {
			m.Lock()         // HL
			defer m.Unlock() // HL

			for i := 0; i < 100; i++ {
				reversed[v] = k
			}
		}(k, v)
	}

	time.Sleep(time.Second)
	fmt.Println(reversed)
}
