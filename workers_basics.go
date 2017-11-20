package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Pool interface {
	Add(string)
	Release() chan bool
	Start()
	Stop() chan os.Signal
}

func init() {
	rand.Seed(time.Now().Unix())
}

type workFunc func(int, string)

type pool struct {
	nWorkers int
	fnc      workFunc

	messageChan chan string

	releaseChan chan bool
	stopChan    chan os.Signal
}

// NEW POOL OMIT
func NewPool(workers int, fnc workFunc) *pool {
	return &pool{
		nWorkers:    workers,
		fnc:         fnc,
		messageChan: make(chan string),
		releaseChan: make(chan bool),
		stopChan:    make(chan os.Signal),
	}
}

// END POOL OMIT

func (p *pool) Add(msg string) {
	p.messageChan <- msg
}

func (p *pool) Start() {
	stop := make(chan bool, p.nWorkers) // HL
	ack := make(chan bool)

	p.startWorkers(stop, ack)
	go p.ackSignal(stop, ack) // HL
}

func (p *pool) ackSignal(stop, ack chan bool) {
	<-p.Stop() // HL
	fmt.Println("Stop requested, shutting down workers...")

	for i := 0; i < p.nWorkers; i++ {
		stop <- true // HL
	}

	for i := 0; i < p.nWorkers; i++ {
		<-ack // HL
	}

	fmt.Println("Stopped workers, shutting down")
	p.Release() <- true // HL
}

func (p *pool) startWorkers(stop, ack chan bool) { // HL
	for i := 0; i < p.nWorkers; i++ {
		go func(v int) {
			for {
				select {
				case msg := <-p.messageChan: // HL
					p.fnc(v, msg)
				case <-stop: // HL
					fmt.Printf("worker[%d] stopping...\n", v)
					ack <- true // HL
					return      // HL
				}
			}
		}(i)
	}
}

func (p *pool) Stop() chan os.Signal {
	return p.stopChan // HL
}

func (p *pool) Release() chan bool {
	return p.releaseChan
}

func main() {
	processor := func(i int, msg string) {
		sleep := rand.Intn(5000)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		fmt.Printf("worker[%d] %s\n", i, msg)
	}

	p := NewPool(4, processor)

	p.Start() // HL

	go func() {
		for {
			// fmt.Println("Adding work to the pool...") // OMIT
			p.Add(fmt.Sprintf("New Message @ %d", time.Now().Unix())) // HL
		}
	}()

	signal.Notify(p.Stop(), syscall.SIGINT) // HL
	<-p.Release()                           // HL
	fmt.Println("Released, exiting")
}
