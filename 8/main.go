package main

import (
	"sync"
	"time"
)

//Не понял, зачем тут семафоры, ведь есть atomic

type WaitG struct {
	counter int
	mu sync.Mutex
}

func (wg *WaitG) Add() {
	wg.mu.Lock()
	defer wg.mu.Unlock()
	wg.counter++
}

func (wg *WaitG) Done() {
	wg.mu.Lock()
	defer wg.mu.Unlock()
	wg.counter--
}

func (wg *WaitG) Wait() {
	for wg.counter > 0{}
}

func main() {
	var wg WaitG
	wg.Add()
	go func () {
		defer wg.Done()
		time.Sleep(time.Second)
	}()
	wg.Wait()
}