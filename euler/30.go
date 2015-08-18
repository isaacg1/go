package main

func fifth(n int) int {
	return n * n * n * n * n
}

func fifth_sum(n int) int {
	sum := 0
	for n > 0 {
		sum += fifth(n % 10)
		n /= 10
	}
	return sum
}

func main() {
	sum := 0
	for i := 10; i < 1e6; i++ {
		res := fifth_sum(i)
		if i == res {
			sum += i
		}
	}
	println(sum)
}
