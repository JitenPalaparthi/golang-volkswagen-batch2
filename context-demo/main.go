package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	parent := context.Background()
	//ctx, cancel := context.WithCancel(parent)
	ctx, cancel := context.WithTimeout(parent, time.Second*5)

	defer cancel()
	// go func() {
	// 	time.Sleep(time.Second * 2)
	// 	cancel()
	// }()
	ch1 := generate(ctx)
	<-receiver(ctx, ch1)
}

func generate(ctx context.Context) chan int {
	ch := make(chan int)
	go func() {
		i := 1
	out:
		for {
			time.Sleep(time.Millisecond * 300)
			select {
			case ch <- i * 1:
			case <-ctx.Done():
				println("context cancelled in generate")
				close(ch)
				break out
			}
			i++
		}
	}()
	return ch
}

func receiver(ctx context.Context, ch chan int) chan struct{} {
	sig := make(chan struct{})
	go func() {
	out:
		for {
			select {
			case v, ok := <-ch:
				fmt.Println(v, ok)
			case <-ctx.Done():
				println("context cancelled in receiver")
				break out
			}
		}
		sig <- struct{}{}
		close(sig)
	}()
	return sig
}

func add(ctx context.Context, a, b int) int {
	return a + b
}
