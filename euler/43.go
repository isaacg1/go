package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	cands := []string{}
	for i := 1; i*17 < 1000; i++ {
		add := true
		str_cand := strconv.Itoa(i * 17)
		for _, digit := range strings.Split(str_cand, "") {
			if strings.Count(str_cand, digit) > 1 {
				add = false
			}
		}
		if add {
			cands = append(cands, str_cand)
		}
	}
	pos := 1000
	for _, div := range []int{13, 11, 7, 5, 3, 2, 1} {
		new_cands := []string{}
		for digit := 0; digit < 10; digit++ {
			str_digit := strconv.Itoa(digit)
			for _, cand := range cands {
				if strings.Count(cand, str_digit) == 0 {
					new_cand := fmt.Sprint(str_digit, cand)
					new_cand_num, err := strconv.Atoi(new_cand[:3])
					if err == nil {
						if new_cand_num%div == 0 {
							new_cands = append(new_cands, new_cand)
						}
					}
				}
			}
		}
		pos *= 10
		cands = new_cands
	}
	sum := 0
	for _, cand := range cands {
		cand_num, _ := strconv.Atoi(cand)
		sum += cand_num
	}
	println(sum)
}
