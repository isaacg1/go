package main

func main() {
	digits := []int{1}
	for i := 0; i < 1000; i++ {
		carry := 0
		for ind, val := range digits {
			if val < 5 {
				digits[ind] = 2*val + carry
				carry = 0
			} else {
				digits[ind] = 2*val + carry - 10
				carry = 1
			}
		}
		if carry > 0 {
			digits = append(digits, carry)
		}
	}
	sum := 0
	for _, val := range digits {
		sum += val
	}
	println(sum)
}
