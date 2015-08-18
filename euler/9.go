package main

func gen_trip(m, n int) (int, int, int) {
	return 2 * m * n, m*m - n*n, m*m + n*n
}

func main() {
	for i := 0; i < 25; i++ {
		for j := 0; j < i; j++ {
			a, b, c := gen_trip(i, j)
			if a+b+c == 1000 {
				println(a, b, c, a*b*c)
			}
		}
	}
}
