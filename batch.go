package main
import (
	"fmt"
	"math"
)
func sumofall(target int) int {
	return target * (target + 1) / 2
}

func classicSum(target int) int {
	res := 0
	i := 0
	//sequence can ignore initializer and stepper
	for ; i < target + 1 ; i++ {
		res += i
	}
	return res
}

func looper(target int) {
	i := 0

	for {
		i ++
		if i == target {
			fmt.Printf("I has reached %v", target)
			break
		}
	}
}

func pow(x, n, lim float64) float64 {
	// if can have initializer!!
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main(){
	fmt.Println(sumofall(3))
	fmt.Println(classicSum(3))
	fmt.Println(pow(3,2,10))
	looper(3)// somehow, this also breaks the main function. don't know how to fix it yet...
}