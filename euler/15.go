package main

const size int = 21

func main() {
	twoD := make([][]int, size)
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			var right int
			var above int
			if x == 0 {
				right = 0
			} else {
				right = twoD[y][x-1]
			}
			if y == 0 {
				above = 0
			} else {
				above = twoD[y-1][x]
			}
			if y == 0 && x == 0 {
				twoD[y] = append(twoD[y], 1)
			} else {
				twoD[y] = append(twoD[y], above+right)
			}
		}
	}
	println(twoD[size-1][size-1])
}
