package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		2.0, 3.0,
	},
	"Google" : Vertex{
		3.4, 5.8,
	},
}

func main() {
	fmt.Println(m)
}
