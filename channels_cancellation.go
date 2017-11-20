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
	timeout := 10 * time.Second // HL
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	resp := make(chan bool)

	for i := 0; i < 3; i++ { // HL
		go expensiveWork(ctx, i, resp) // HL
	} // HL

	fmt.Println(<-resp) // HL
	cancel()            // HL

	// for demo purpose only, ensure that output is printed
	time.Sleep(time.Second)
	fmt.Println("Done!")
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
		fmt.Printf("Routine %d: context is cancelled\n", i)
		ret <- false
	}
}
