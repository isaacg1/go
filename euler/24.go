package main

import (
	"fmt"
	"strings"
)

func fact(n int) int {
	prod := 1
	for i := 1; i <= n; i++ {
		prod *= i
	}
	return prod
}

func nth_lexi(n int, size int) string {
	unused := ""
	for i := 0; i < size; i++ {
		unused = fmt.Sprint(unused, i)
	}
	res := ""
	for pos := size - 1; pos >= 0; pos-- {
		group_size := fact(pos)
		unused_pos := n / group_size
		n %= group_size
		unused_val := unused[unused_pos : unused_pos+1]
		res = fmt.Sprint(res, unused_val)
		unused = strings.Join(strings.Split(unused, unused_val), "")
	}
	return res
}

func main() {
	println(nth_lexi(1000000-1, 10))
}
