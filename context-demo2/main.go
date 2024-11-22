package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGKILL, syscall.SIGTERM)
	defer cancel()
	ch, sig := generate(ctx)
	go func() {
		for v := range ch {
			println(v)
		}
	}()
	<-sig
}

func generate(ctx context.Context) (chan int, chan struct{}) {
	ch := make(chan int)
	sig := make(chan struct{})
	go func() {
		i := 1
	out:
		for {
			time.Sleep(time.Millisecond * 300)
			select {
			case ch <- i * 1:
			case <-ctx.Done():
				println("context cancelled in generate")
				sig <- struct{}{}
				close(ch)
				close(sig)
				break out
			}
			i++
		}
	}()
	return ch, sig
}
