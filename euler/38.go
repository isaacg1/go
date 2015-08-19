package main

import (
	"math"
	"strconv"
	"strings"
)

func concat(nums []int) string {
	str_vals := make([]string, len(nums))
	for i, num := range nums {
		str_vals[i] = strconv.Itoa(num)
	}
	return strings.Join(str_vals, "")
}

func is_pan(str string) bool {
	if len(str) != 9 {
		return false
	}
	for i := 1; i < 10; i++ {
		if strings.Count(str, strconv.Itoa(i)) != 1 {
			return false
		}
	}
	return true
}

func main() {
	max := ""
	for n := 2; n < 10; n++ {
		for i := 0; i < int(math.Pow10(int(math.Ceil(10/float64(n))))); i++ {
			prods := []int{}
			for j := 1; j <= n; j++ {
				prods = append(prods, j*i)
			}
			str_prod := concat(prods)
			if is_pan(str_prod) {
				println(i, n, str_prod)
				if max < str_prod {
					max = str_prod
				}
			}
		}
	}
	println(max)
}
