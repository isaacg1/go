package main

import (
	"io/ioutil"
	"sort"
	"strings"
)

func value(name string) int {
	sum := 0
	for _, val := range name {
		sum += int(val) - 64
	}
	return sum
}

func main() {
	file_bytes, err := ioutil.ReadFile("p022_names.txt")
	if err != nil {
		println(err.Error())
	} else {
		names := string(file_bytes)
		names = strings.Replace(names, "\"", "", -1)
		names_list := strings.Split(names, ",")
		sort.Strings(names_list)
		sum := 0
		for i, name := range names_list {
			sum += (i + 1) * value(name)
		}
		println(sum)
	}
}
