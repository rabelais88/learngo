package main

import (
	//"fmt"

	"strings"

	"golang.org/x/tour/wc"
)

//intially, the tour recommended using string.Fields()
//but I haven't used it - sungryeol

func WordCount(s string) map[string]int {
	words := strings.Split(s, " ")
	counts := make(map[string]int)
	for _, element := range words {
		if _, ok := counts[element]; ok {
			counts[element]++
		} else {
			counts[element] = 1
		}
	}
	return counts
}

func main() {
	wc.Test(WordCount)
}
