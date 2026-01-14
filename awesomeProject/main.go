package main

import (
	"fmt"
	"regexp"
	"sync"
	"time"
)

// Доработайте тип Worker
// с помощью sync.Cond.

type Worker struct {
	c     sync.Cond
	ready bool
}

func NewWorker() *Worker {
	w := &Worker{}
	w.c = *sync.NewCond(&sync.Mutex{})
	return w
}

func (w *Worker) setReady() {
	w.c.L.Lock()
	time.Sleep(2 * time.Second)
	w.ready = true
	w.c.Broadcast()
	w.c.L.Unlock()
}

func (w *Worker) CheckReady() bool {
	return w.ready
}

func (w *Worker) WaitReady() {
	w.c.L.Lock()
	for !w.CheckReady() {
		w.c.Wait()
	}
	w.c.L.Unlock()
}

func main1() {
	w := NewWorker()

	go func() { // имитация инициализации
		time.Sleep(time.Second)
		fmt.Println("I'm ready")
		w.setReady()
	}()

	w.WaitReady()

	fmt.Println("Ready to use Worker")
}

func main() {
	regexp.MustCompile(`[!@#$\n-.,]`)
}
