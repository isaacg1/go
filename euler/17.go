package main

func eng_len(n int) int {
	count := 0
	if n >= 100 {
		count += eng_len(n / 100)
		count += len("hundred")
		if n%100 > 0 {
			count += len("and")
		}
	}
	rem := n % 100
	if rem < 20 {
		special_cases := []int{0, 3, 3, 5, 4, 4, 3, 5, 5, 4, 3, 6, 6, 8, 8, 7, 7, 9, 8, 8}
		count += special_cases[rem]
	} else {
		tens := []int{6, 6, 5, 5, 5, 7, 6, 6}
		count += tens[rem/10-2]
		if rem%10 > 0 {
			count += eng_len(n % 10)
		}
	}
	return count
}

func main() {
	sum := len("onethousand")
	for i := 1; i < 1000; i++ {
		sum += eng_len(i)
	}
	println(sum)
}
