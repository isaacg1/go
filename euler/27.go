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

func main() {
	most_primes := 0
	best_prod := 0
	for a := -999; a < 1000; a++ {
		for b := -999; b < 1000; b++ {
			n := 0
			for ; is_prime(n*n + a*n + b); n++ {
			}
			if n > most_primes {
				most_primes = n
				best_prod = a * b
			}
		}
	}
	println(best_prod)
}
