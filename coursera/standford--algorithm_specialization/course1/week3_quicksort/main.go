package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Comparisons if the first element is pivot

// 12 2 10 1 = 3 - 0
// 1 2 10 12 = 2 - 0
// 1 2 10 12

// 4 2 10 11 12 = 4
// 2 4 10 11 12 = 3
// 2 = 0
// 10 11 = 1
// 11 = 0

func main() {
	var arr []int
	var err error
	var m int

	arr, err = readFromFile("./numbers.txt")
	if err != nil {
		fmt.Println("Errrror!")
		return
	}
	// arr = []int{12, 2, 10, 11, 4}
	m = QuickSort(arr, 0, len(arr)-1, "first")
	fmt.Println(m)
	fmt.Println(arr[len(arr)-10:])

	arr, err = readFromFile("./numbers.txt")
	if err != nil {
		fmt.Println("Errrror!")
		return
	}
	// arr = []int{12, 2, 10, 11, 4}
	m = QuickSort(arr, 0, len(arr)-1, "last")
	fmt.Println(m)
	fmt.Println(arr[len(arr)-10:])

	arr, err = readFromFile("./numbers.txt")
	if err != nil {
		fmt.Println("Errrror!")
		return
	}
	// arr = []int{12, 2, 10, 11, 4}
	m = QuickSort(arr, 0, len(arr)-1, "medianOf3")
	fmt.Println(m)
	fmt.Println(arr[len(arr)-10:])
}

func QuickSort(arr []int, left int, right int, mode string) int {
	if left >= right {
		return 0
	}

	if mode == "last" {
		arr[right], arr[left] = arr[left], arr[right]
	} else if mode == "medianOf3" {
		middle := (right-left)/2 + left
		if (arr[right] <= arr[left] && arr[right] >= arr[middle]) ||
			(arr[right] >= arr[left] && arr[right] <= arr[middle]) {
			arr[right], arr[left] = arr[left], arr[right]
		} else if (arr[middle] <= arr[left] && arr[middle] >= arr[right]) ||
			(arr[middle] >= arr[left] && arr[middle] <= arr[right]) {
			arr[middle], arr[left] = arr[left], arr[middle]
		}
	}

	pivot := left
	i := left + 1
	for j := i; j <= right; j++ {
		if arr[j] < arr[pivot] {
			arr[j], arr[i] = arr[i], arr[j]
			i++
		}
	}

	arr[pivot], arr[i-1] = arr[i-1], arr[pivot]
	pivot = i - 1

	// fmt.Println(arr, right-left)

	lComp := QuickSort(arr, left, pivot-1, mode)
	rComp := QuickSort(arr, pivot+1, right, mode)

	return lComp + rComp + (right - left)
}

func readFromFile(filename string) ([]int, error) {
	arr := []int{}
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		arr = append(arr, num)
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return arr, nil
}
