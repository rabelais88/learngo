package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	var m2 = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}

	m3 := make(map[string]int)
	m3["answer"] = 42
	m3["answer"] = 48
	m3["wrong"] = 2
	delete(m3, "answer")

	fmt.Println(m)
	fmt.Println(m2)
	fmt.Println(m3)
}
