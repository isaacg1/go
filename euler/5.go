package main

func make_div_by(n int, div int) int {
	max_divisor := 0
	for i := 1; i <= div; i++ {
		if (n%i == 0) && (div%i == 0) {
			max_divisor = i
		}
	}
	return (div / max_divisor) * n
}

func main() {
	lcm := 1
	for i := 1; i <= 20; i++ {
		lcm = make_div_by(lcm, i)
	}
	println(lcm)
}
