package main

import "math"

func is_prime(p int) bool {
	if p < 2 {
		return false
	}
	for i := 2; i*i <= p; i++ {
		if p%i == 0 {
			return false
		}
	}
	return true
}

func circular(p int) bool {
	digits := int(math.Log10(float64(p))) + 1
	pow := 1
	for i := 0; i < digits-1; i++ {
		pow *= 10
	}
	for i := 0; i < digits; i++ {
		if !is_prime(p) {
			return false
		}
		p = (p%10)*pow + p/10
	}
	return true
}

const upper int = 1e6

func main() {
	count := 0
	for i := 1; i < upper; i++ {
		if circular(i) {
			count++
		}
	}
	println(count)
}
