package main

func square(n int) int {
	return n * n
}

func main() {
	sum := 1
	for spiral := 0; spiral < 1001/2; spiral++ {
		start := square(2*spiral + 1)
		end := square(2*spiral + 3)
		side := (end - start) / 4
		for i := 1; i <= 4; i++ {
			sum += start + i*side
		}
	}
	println(sum)
}
