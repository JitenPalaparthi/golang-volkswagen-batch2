package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	var ch chan int     // declared a channel
	ch = make(chan int) // instntiate a channel
	wg.Add(1)
	go func() {
		v := <-ch // receiving the data
		fmt.Println("seems the block is released just now")
		println(v)
		wg.Done()
	}()
	fmt.Println("i come here but waiting")
	ch <- 100
	wg.Wait()
}

// Do not communicate by sharing memory; instead, share memory by communicating.
// 1. Channels -> used for communication between goroutines
// 2. Channel-> kind of a conduit
// 3. Channels: buffered and unbuffered
// 4. There is a sender and receiver (generally different goroutines)
// 		sender is blocked until the receiver receives the data
// 		receiver is blocked until the sender sends the data
// 5. To send data ch <- 100
// 6. To receive data <-ch
