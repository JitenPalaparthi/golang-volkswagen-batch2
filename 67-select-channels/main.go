package main

import "runtime"

func main() {
	ch1 := Generate(10, 20)
	ch2 := Generate(100, 120)
	<-Receiver(ch1, ch2)
}

func Generate(from, to int) chan int {
	max := max(from, to)
	min := min(from, to)
	ch := make(chan int)
	go func() {
		for i := min; i <= max; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func Receiver(ch1, ch2 <-chan int) chan struct{} {
	sig := make(chan struct{})
	done1, done2 := false, false
	go func() {
		for {
			if done1 && done2 {
				sig <- struct{}{}
				runtime.Goexit()
			}
			select {
			case v1, ok1 := <-ch1:
				if ok1 {
					println(v1)
				} else if ok1 == false && done1 == false {
					done1 = true
				}
			case v2, ok2 := <-ch2:
				if ok2 {
					println(v2)
				} else if ok2 == false && done2 == false {
					done2 = true
				}
			}

		}
	}()

	return sig
}

// Generator
// Future
// Pipelines
// Fanin, Fanout
