package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ch := Generator(2) // from 1-20
	sqch, cubech := Receive1(ch)

	map1 := make(map[int]any)
	slice1 := make([]int, 0)
	ch2 := Generator(20) // from 1-20
	go func() {
		for v := range ch2 {
			map1[v] = fmt.Sprintf("map-value-go1-%d", v)
			slice1 = append(slice1, v)
		}
	}()
	go func() {
		for v := range ch2 {
			map1[v] = fmt.Sprintf("map-value-go2-%d", v)
			slice1 = append(slice1, v)
		}
	}()
	go func() {
		for v := range ch2 {
			map1[v] = fmt.Sprintf("map-value-go3-%d", v)
			slice1 = append(slice1, v)
		}
	}()

	<-Receive2(sqch, cubech)
	time.Sleep(time.Second * 1)
	fmt.Println(map1)
	fmt.Println("Slice1:", slice1)

}

func Generator(num uint) chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= int(num); i++ {
			time.Sleep(time.Millisecond * 100)
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func Receive1(ch <-chan int) (chan int, chan int) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for v := range ch {
			ch1 <- v * v
			ch2 <- v * v * v
		}
		close(ch1)
		close(ch2)
	}()

	return ch1, ch2
}

func Receive2(ch1 <-chan int, ch2 <-chan int) chan struct{} {
	sig := make(chan struct{})
	go func() {
		done1, done2 := false, false
		for {
			if done1 && done2 {
				sig <- struct{}{}
				close(sig)
				runtime.Goexit()
			}
			select {
			case v1, ok1 := <-ch1:
				if ok1 {
					println("Square:", v1)
				} else if ok1 == false && done1 == false {
					done1 = true
				}

			case v2, ok2 := <-ch2:
				if ok2 {
					println("Cute:", v2)
				} else if ok2 == false && done2 == false {
					done2 = true
				}
			}
		}

	}()

	return sig
}
