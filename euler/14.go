package main

func collatz(n int) int {
	count := 0
	for n > 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		count++
	}
	return count + 1
}
func main() {
	max_steps := 0
	best_start := 1
	for i := 1; i < 1e6; i++ {
		steps := collatz(i)
		if steps > max_steps {
			max_steps = steps
			best_start = i
		}
	}
	println(best_start, max_steps)
}
