package main

func main() {

	ch := make(chan int)
	count := make(chan int)

	go sender(ch)
	go receiver(ch, count)

	println(<-count) // print count and also block main

}

func sender(send chan<- int) {
	for i := 1; i <= 10; i++ {
		send <- i * i
	}
	close(send)
}

func receiver(receive <-chan int, count chan<- int) {
	i := 0
outer:
	for {
		select {
		case v, ok := <-receive:
			if ok {
				i++
				println(v, ok)
			} else {
				break outer
			}
		}
	}
	// for v := range receive {
	// 	println(v)
	// 	i++
	// }

	// v1, ok := <-receive
	// fmt.Println("Closed?", v1, ok)
	// fmt.Println("Closed?", v1, ok)
	count <- i
}
