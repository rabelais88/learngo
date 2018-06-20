package main
import (
	"fmt"
	"math/rand"
	"math/cmplx"
	"math"
	"time"
)
func main(){
	var i = 0
	power := math.Pow(10, 1.0/3.0)
	const (
		PI = math.Pi
		MAXNUMBER = 256 
	)

	var (
		tf			bool = false
		maxInt	uint64 = 1<< 64 - 1
		z				complex128 = cmplx.Sqrt(-5 + 12i)
		f				float64 //values can remain undefined
	)
	f = 123.000
	fmt.Println(f)
	fmt.Println(i)
	fmt.Println("randomized number:", randomer(MAXNUMBER))
	fmt.Printf("cube root of 10 : %v \n", power)
	fmt.Printf("PI : %v \n", PI)
	fmt.Printf("Type: %T Value: %v \n", tf, tf)
	fmt.Printf("Type: %T Value: %v \n", maxInt, maxInt)
	fmt.Printf("Type: %T Value: %v \n", z, z)
}

func randomer(max int) int{
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max)
}