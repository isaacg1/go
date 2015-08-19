package main

func is_prime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func prime_plus_twice_square(n int) bool {
	for i := 0; i*i*2 < n; i++ {
		if is_prime(n - i*i*2) {
			return true
		}
	}
	return false
}

func main() {
	var n int
	for n = 3; prime_plus_twice_square(n); n += 2 {
	}
	println(n)
}
