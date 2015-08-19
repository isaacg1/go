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
	left := []int{2, 3, 5, 7}
	ends := []int{1, 3, 7, 9}
	sum := 0
	for ten_pow := 10; len(left) > 0; ten_pow *= 10 {
		new_left := []int{}
		for _, prime := range left {
			for _, end := range ends {
				cand := prime*10 + end
				if is_prime(cand) {
					new_left = append(new_left, cand)
					tmp_pow := ten_pow
					tmp_cand := cand
					for tmp_cand > 10 {
						tmp_cand %= tmp_pow
						if !is_prime(tmp_cand) {
							break
						}
						tmp_pow /= 10
					}
					if tmp_pow == 1 {
						println(cand)
						sum += cand
					}
				}
			}
		}
		left = new_left
	}
	println(sum)
}
