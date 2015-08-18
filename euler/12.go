package main

func divisors(n int) int {
	count := 0
	i := 1
	for ; i*i < n; i++ {
		if n%i == 0 {
			count++
		}
	}
	count *= 2
	if i*i == n {
		count++
	}
	return count
}

func triang(n int) int {
	return (n * (n + 1)) / 2
}

func main() {
	i := 0
	for divisors(triang(i)) < 500 {
		i++
	}
	println(triang(i))
}
