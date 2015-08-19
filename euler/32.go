package main

import (
	"fmt"
	"sort"
	"strings"
)

func is_pan(n int) bool {
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			concat := strings.Split(fmt.Sprintf("%v%v%v", i, n/i, n), "")
			sort.Strings(concat)
			if strings.Join(concat, "") == "123456789" {
				return true
			}
		}
	}
	return false
}

func main() {
	sum := 0
	for i := 0; i < 1e4; i++ {
		if is_pan(i) {
			sum += i
		}
	}
	println(sum)
}
