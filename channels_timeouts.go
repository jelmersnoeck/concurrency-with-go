package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	timeout := 200 * time.Millisecond // HL
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	resp := make(chan bool)
	go expensiveWork(ctx, 0, resp) // HL
	fmt.Println(<-resp)
}

func expensiveWork(ctx context.Context, i int, ret chan bool) {
	c := make(chan bool)
	go func() {
		r := time.Duration(rand.Intn(500)) * time.Millisecond // HL
		fmt.Printf("Routine %d: sleeping %s\n", i, r)
		time.Sleep(r)
		c <- true
	}()

	select {
	case <-c:
		fmt.Printf("Routine %d: got value\n", i)
		ret <- true
	case <-ctx.Done(): // HL
		fmt.Printf("Routine %d: context timeout\n", i)
		ret <- false
	}
}
