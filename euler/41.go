package main

import (
	"strconv"
	"strings"
)

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

func is_pan(str string) bool {
	for i := 1; i <= len(str); i++ {
		if strings.Count(str, strconv.Itoa(i)) != 1 {
			return false
		}
	}
	return true
}

func main() {
	for i := 7660001; i > 7650000; i -= 2 {
		if is_prime(i) && is_pan(strconv.Itoa(i)) {
			println(i)
			break
		}
	}
}
