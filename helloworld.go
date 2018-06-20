package main

import (
	"fmt"
	"strconv"
)

func getSum(a, b int) string{
	//because sum of a and b is integer, res is declared as int
	res := a + b
	return strconv.Itoa(a) + " + " + strconv.Itoa(b) +
	" = " + strconv.Itoa(res)
}

func sumAndsub(a, b int) (sum, sub int){
	sum = a + b
	sub = a - b
	return
}

func main(){

	// all exported functions must start with uppercase
	fmt.Println("hello world!")
	fmt.Println(getSum(1,2))
	fmt.Println(sumAndsub(1,2))
}