package main

const upper int = 1e6

func fact(n int) int {
	prod := 1
	for i := 1; i <= n; i++ {
		prod *= i
	}
	return prod
}

func fact_sum(n int) int {
	sum := 0
	for n > 0 {
		sum += fact(n % 10)
		n /= 10
	}
	return sum
}

func main() {
	sum := 0
	for i := 10; i < upper; i++ {
		if fact_sum(i) == i {
			println(i)
			sum += i
		}
	}
	println(sum)
}
