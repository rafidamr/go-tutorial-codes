package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var mtx = make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		t := make([]uint8, dx)
		mtx[i] = t
	}

	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			mtx[y][x] = uint8((x + y) / 2)
		}
	}
	return mtx
}

func main() {
	pic.Show(Pic)
}
