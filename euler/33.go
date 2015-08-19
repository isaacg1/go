package main

import "fmt"

func cancels(n int) bool {
	if n%10 == 0 || n/100 == 0 {
		return false
	}
	num1, denom1 := lowest_terms(n/10, n%100)
	num2, denom2 := lowest_terms(n/100, n%10)
	return num1 == num2 && denom1 == denom2 && num1 < denom1
}

func lowest_terms(num, denom int) (int, int) {
	for i := 1; i <= denom; i++ {
		if (num*i)%denom == 0 {
			return (num * i) / denom, i
		}
	}
	return num, denom
}

func main() {
	num := 1
	denom := 1
	for i := 0; i < 1000; i++ {
		if cancels(i) {
			fmt.Printf("%v/%v = %v/%v\n", i/10, i%100, i/100, i%10)
			num *= i / 100
			denom *= i % 10
		}
	}
	low_num, low_denom := lowest_terms(num, denom)
	fmt.Printf("%v/%v\n", low_num, low_denom)
}
