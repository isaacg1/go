package main

import (
	"bigint"
)

func main() {
	num1 := bigint.BigInt{[]int{1}}
	num2 := num1.Copy()
	var index int
	for index = 2; len(num1.String()) < 1000; index++ {
		tmp := num1.Copy()
		num1.Add(num2)
		num2 = tmp
	}
	println(index)
}
