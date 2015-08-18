package main

import (
	"bigint"
)

func main() {
	prod := bigint.BigInt{[]int{1}}
	for i := 0; i < 1000; i++ {
		prod.Mult(2)
	}
	println(prod.DigitSum())
}
