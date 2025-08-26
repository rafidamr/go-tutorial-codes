package main

import (
	"fmt"
	"strconv"
)

func sliceFunc() {
	// a reference to an array with capacity of 4
	slc := []int{1, 2, 3, 4}
	printSlice(slc)

	slc = slc[:2]
	printSlice(slc)

	slc = slc[:4]
	printSlice(slc)

	slc = slc[2:]
	printSlice(slc)

	expand_outofrange(slc)

	slc2 := make([]int, 0, 10)
	printSlice(slc2)
	slc2 = slc2[:cap(slc2)]
	printSlice(slc2)
	slc2 = append(slc2, 1, 2, 3)
	printSlice(slc2)

	for idx := range slc2 {
		fmt.Print(strconv.Itoa(idx) + " ")
	}
}

func expand_outofrange(slc []int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	slc = slc[:4]
	printSlice(slc)
}

// creating a matrix
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

func printSlice(slc []int) {
	fmt.Printf("len=%v cap=%v slc=%v\n", len(slc), cap(slc), slc)
}
