package main

const SQRT int = 1415
const SIZE int = 2000000

func main() {
	is_composite := make([]bool, SIZE)
	for prime := 2; prime <= SQRT; prime++ {
		if !is_composite[prime] {
			for mult := prime * 2; mult < SIZE; mult += prime {
				is_composite[mult] = true
			}
		}
	}
	sum := 0
	for i := 2; i < SIZE; i++ {
		if !is_composite[i] {
			sum += i
		}
	}
	println(sum)
}
