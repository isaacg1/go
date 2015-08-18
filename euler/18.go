package main

import (
	"strconv"
	"strings"
)

const data string = `75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23`

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func main() {
	lines := strings.Split(data, "\n")
	best_paths := make([][]int, len(lines))
	first_num, _ := strconv.Atoi(lines[0])
	best_paths[0] = append(best_paths[0], first_num)
	for y, line := range lines {
		if y > 0 {
			split_line := strings.Fields(line)
			for x, str_of_num := range split_line {
				num, _ := strconv.Atoi(str_of_num)
				var best_path int
				if x == 0 {
					best_path = best_paths[y-1][0]
				} else if x == len(split_line)-1 {
					best_path = best_paths[y-1][x-1]
				} else {
					best_path = max(best_paths[y-1][x-1], best_paths[y-1][x])
				}
				best_paths[y] = append(best_paths[y], best_path+num)
			}
		}
	}
	max := 0
	for _, path := range best_paths[len(best_paths)-1] {
		if path > max {
			max = path
		}
	}
	println(max)
}
