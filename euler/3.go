package main

import (
	"math"
)

func largest_prime_fact(n int) int {
	i := 2
	for i <= int(math.Sqrt(float64(n))) {
		if n%i == 0 {
			n /= i
		} else {
			i++
		}
	}
	return n
}

func main() {
	println(largest_prime_fact(600851475143))
}
