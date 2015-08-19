package main

import (
	"io/ioutil"
	"math"
	"strings"
)

func value(word string) int {
	sum := 0
	for _, val := range word {
		sum += int(val) - 64
	}
	return sum
}

func is_triang(n int) bool {
	t := int(math.Sqrt(float64(n * 2)))
	return (t*(t+1))/2 == n
}

func main() {
	file_bytes, err := ioutil.ReadFile("p042_words.txt")
	if err != nil {
		println(err.Error())
	} else {
		words := string(file_bytes)
		words = strings.Replace(words, "\"", "", -1)
		words_list := strings.Split(words, ",")

		count := 0
		for _, word := range words_list {
			val := value(word)
			if is_triang(val) {
				count += 1
			}
		}
		println(count)
	}
}
