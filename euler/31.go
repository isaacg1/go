package main

func main() {
	coins := []int{1, 2, 5, 10, 20, 50, 100, 200}
	combs := make([]int, 201)
	combs[0] = 1
	for _, coin := range coins {
		for i := 0; i <= 200; i++ {
			if i+coin <= 200 {
				combs[i+coin] += combs[i]
			}
		}
	}
	println(combs[200])
}
