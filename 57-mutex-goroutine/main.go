package main

import (
	"sync"
	"time"
)

var (
	global int = 1
	wg     *sync.WaitGroup
	mu     *sync.Mutex
)

// type MoneyCounter struct {
// 	Counter int
// 	mu      *sync.Mutex
// }

// func (mc *MoneyCounter) Increment() {
// 	mc.mu.Lock()
// 	mc.Counter++
// 	mc.mu.Unlock()

// }
// func (mc *MoneyCounter) Decrement() {
// 	mc.mu.Lock()
// 	mc.Counter--
// 	mc.mu.Unlock()
// }

func init() {
	wg = new(sync.WaitGroup)
	mu = new(sync.Mutex)
}

func main() {
	wg.Add(1)
	go func() {
		for i := 1; i <= 1000; i++ {
			increment(mu)
		}
		wg.Done()
	}()
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		for i := 1; i <= 100; i++ {
			decrement(mu)
		}
		wg.Done()
	}(wg)
	wg.Wait()
	println("Global:", global)
}

func increment(mu *sync.Mutex) {
	mu.Lock()
	time.Sleep(time.Millisecond * 100)
	global++
	mu.Unlock()
}

func decrement(mu *sync.Mutex) {
	mu.Lock()
	global--
	mu.Unlock()
}
