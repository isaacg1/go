package main

func div_sum(n int) int {
	div := 1
	sum := 0
	for div*div < n {
		if n%div == 0 {
			sum += div
			sum += n / div
		}
		div++
	}
	if div*div == n {
		sum += div
	}
	return sum
}

func main() {
	sum := 0
	for i := 1; i < 10000; i++ {
		pair := div_sum(i) - i
		if i == div_sum(pair)-pair && i != pair {
			sum += i
		}
	}
	println(sum)
}
