package main

import (
	"bigint"
)

func main() {
	fact := bigint.BigInt{[]int{1}}
	for i := 1; i <= 100; i++ {
		fact.Mult(i)
	}
	println(fact.String())
	println(fact.DigitSum())
}
