//defer is like delayed decorator
//they run after the surrouned function
package main
import "fmt"
func main() {
	defer fmt.Println("world")
	hello := func() { fmt.Println("hello") }
	hello() //go singleton

	func() { fmt.Println("this is singleton") }()

	//defer can also be stacked
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}