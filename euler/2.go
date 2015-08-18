package main

func fib_stream(ch chan int) {
	prev, curr := 1, 1
	for curr < 4e6 {
		ch <- curr
		prev, curr = curr, prev+curr
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	sum := 0
	go fib_stream(ch)
	for fib := range ch {
		if fib%2 == 0 {
			sum += fib
		}
	}
	println(sum)
}
