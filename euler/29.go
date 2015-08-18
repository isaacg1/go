package main

const upper int = 100

func main() {
	is_not_base := make([]bool, upper+1)
	powers := 0
	for a := 2; a <= upper; a++ {
		if !is_not_base[a] {
			a_powers := make(map[int]bool)
			for prod, pow := a, 1; prod <= upper; prod, pow = prod*a, pow+1 {
				is_not_base[prod] = true
				for i := 2; i <= upper; i++ {
					a_powers[pow*i] = true
				}
			}
			powers += len(a_powers)
		}
	}
	println(powers)
}
