package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	//golang uses bracket for array because that's mathematical norm
	primes := [6]int{2, 3, 5, 7, 11, 13} 
	fmt.Println(primes)
	var sliced []int = primes[1:4]
	fmt.Println(sliced)

	//and please notice unlike javascript, slice is a converted pointer
	//so if sliced value is modified, then the original value also modifies


	//creating array with struct is also possible
	//and not always need to declare the length
	//but this is not a dynamic array yet!
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)


	var dynamic []int //this nil slice, which is a dynamic array essentially
	checkarray(dynamic)
	dynamic = append(dynamic, 1)
	dynamic = append(dynamic, 5)
	checkarray(dynamic)

	emptyarr := make([]int, 0, 5)
	checkarray(emptyarr)

}

func checkarray(target []int) {
	fmt.Printf("length: %v, cap: %v, %v\n", len(target), cap(target), target)
	return
}
