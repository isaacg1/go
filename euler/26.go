package main

func cycle_len(n int) int {
	rems := make([]int, n)
	dec := 1
	rem := 1
	for rems[rem] == 0 {
		rems[rem] = dec
		rem *= 10
		rem %= n
		dec++
	}
	return dec - rems[rem]
}

func main() {
	max_cycle := 0
	best_div := 0
	for i := 2; i <= 1000; i++ {
		cycle := cycle_len(i)
		if cycle > max_cycle {
			max_cycle = cycle
			best_div = i
		}
	}
	println(best_div)
}
