package main

import "math"

func pent(n int) int {
	return n * (3*n - 1) / 2
}

func is_pent(n int) bool {
	root := int(math.Sqrt((float64(n)*2)/3)) + 1
	return pent(root) == n
}

const upper int = 2500

func main() {
	for i := 1; i < upper; i++ {
		for j := 1; j < upper; j++ {
			sum := pent(i) + pent(j)
			diff := pent(i) - pent(j)
			if is_pent(sum) && is_pent(diff) {
				println(diff)
			}
		}
	}
}
