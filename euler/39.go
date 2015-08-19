package main

func gen_trip(m, n int) (int, int, int) {
	return 2 * m * n, m*m - n*n, m*m + n*n
}

func main() {
	perims := make([]int, 1001)

	for m := 0; m < 25; m++ {
		for n := m%2 + 1; n < m; n += 2 {
			a, b, c := gen_trip(m, n)
			perim := a + b + c
			for i := 1; i*perim <= 1000; i++ {
				perims[i*perim]++
			}
		}
	}
	max := 0
	best := 0
	for ind, val := range perims {
		if val > max {
			max = val
			best = ind
		}
	}
	println(best, max)
}
