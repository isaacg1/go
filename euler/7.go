package main

const SQRT int = 1000
const SIZE int = SQRT * SQRT

func main() {
	is_composite := make([]bool, SIZE)
	for prime := 2; prime <= SQRT; prime++ {
		if !is_composite[prime] {
			for mult := prime * 2; mult < SIZE; mult += prime {
				is_composite[mult] = true
			}
		}
	}
	count := 0
	for i := 2; i < SIZE; i++ {
		if !is_composite[i] {
			count++
			if count == 10001 {
				println(i)
			}
		}
	}
	println(count)
}
