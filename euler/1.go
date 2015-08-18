package main

func div_by_3_or_5(n int) bool {
	return (n%3 == 0) || (n%5 == 0)
}

func main() {
	sum := 0
	for i := 0; i < 1000; i++ {
		if div_by_3_or_5(i) {
			sum += i
		}
	}
	println(sum)
}
