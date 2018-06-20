package main

import (
	"fmt"
	"math/rand"
	"time"
)

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func randomer(max int) int {
	return rand.Intn(max)
}

func main() {
	//this seed randomizer should be declared only once in the whole program
	rand.Seed(time.Now().UnixNano())

	// in go, range substitutes for ~ in and map
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	//this is same as new Array(20).fill(0).map(el=>{})...
	myarr := make([]int, 20)
	const MAXVAL = 20
	for i := range myarr {
		myarr[i] = randomer(20)
	}

	//when index is not necessary and only value is needed,
	//just use _
	for _, value := range myarr {
		fmt.Printf("%d\n", value)
	}
}
