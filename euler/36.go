package main

func base(n int, base int) []int {
	if n == 0 {
		return make([]int, 1)
	}
	conv := []int{}
	for n > 0 {
		conv = append(conv, n%base)
		n /= base
	}
	return conv
}

func is_palin(digits []int) bool {
	for i := 0; i*2 < len(digits); i++ {
		if digits[i] != digits[len(digits)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	sum := 0
	for i := 0; i < 1e6; i++ {
		if is_palin(base(i, 2)) && is_palin(base(i, 10)) {
			sum += i
			println(i)
		}
	}
	println(sum)
}
