package main

func div_sum(n int) int {
	div := 1
	sum := 0
	for div*div < n {
		if n%div == 0 {
			sum += div
			sum += n / div
		}
		div++
	}
	if div*div == n {
		sum += div
	}
	return sum
}

const limit int = 28124

func main() {
	abundant := []int{}
	for i := 0; i < limit; i++ {
		if div_sum(i) > 2*i {
			abundant = append(abundant, i)
		}
	}
	is_sum := make([]bool, limit)
	for _, ab1 := range abundant {
		for _, ab2 := range abundant {
			if ab1+ab2 < limit {
				is_sum[ab1+ab2] = true
			}
		}
	}
	sum := 0
	for ind, flag := range is_sum {
		if !flag {
			sum += ind
		}
	}
	println(sum)
}
