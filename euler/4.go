package main

import (
	"fmt"
)

func is_palindrome(n int) bool {
	str_n := fmt.Sprint(n)
	for i := 0; i < len(str_n)/2; i++ {
		if str_n[i] != str_n[len(str_n)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	max := 0
	for i := 999; i >= 100; i-- {
		for j := 999; j >= 100; j-- {
			cand := i * j
			if cand > max && is_palindrome(cand) {
				max = cand
			}
		}
	}
	println(max)
}
