package main

import "math"

func is_triang(n int) bool {
	t := int(math.Sqrt(float64(n * 2)))
	return (t*(t+1))/2 == n
}

func is_pent(n int) bool {
	root := int(math.Sqrt((float64(n)*2)/3)) + 1
	return (root*(3*root-1))/2 == n
}

func hex(n int) int {
	return n * (2*n - 1)
}

func main() {
	n := 144
	for ; !(is_triang(hex(n)) && is_pent(hex(n))); n++ {
	}
	println(hex(n))
}
