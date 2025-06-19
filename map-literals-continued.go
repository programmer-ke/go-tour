package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell": {1.0, 2.0},
	"Goog": {2.3, 3.2},
}

func main() {
	fmt.Println(m)
}
