package main

//import "golang.org/x/tour/pic"
import "fmt"

func Pic(dx, dy int) [][]uint8 {
	var pic = make([][]uint8, dy)
	for y := range pic {
		row := make([]uint8, dx)
		for x := range row {
			row[x] = uint8(x * y)
		}
		pic[y] = row
	}
	return pic

}

func main() {
	fmt.Println(Pic(3, 3))
}
