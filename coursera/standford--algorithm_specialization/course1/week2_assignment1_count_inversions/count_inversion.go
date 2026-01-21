package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// arr := []int{1, 1, 200, 10, 30, 40}
	arr, err := readFromFile("./numbers.txt")
	if err != nil {
		fmt.Println("Errror!")
		return
	}
	_, inversionCount := mergeSort(arr)
	fmt.Println(inversionCount)
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

func mergeSort(arr []int) ([]int, int) {
	_len := len(arr)
	if _len == 1 {
		return arr, 0
	}
	left, lCount := mergeSort(arr[:(_len / 2)])
	right, rCount := mergeSort(arr[(_len / 2):])
	newArr, splitCount := merge(left, right)
	return newArr, splitCount + rCount + lCount
}

// Piggybacking the merge sort to count the number of split inversion
func merge(left []int, right []int) ([]int, int) {
	inversionCount := 0
	newArr := make([]int, len(left)+len(right))
	i := 0
	j := 0
	k := 0
	for k < len(left)+len(right) {
		if left[i] <= right[j] {
			newArr[k] = left[i]
			i++
		} else {
			newArr[k] = right[j]
			j++
			inversionCount += len(left) - i
		}
		k++
		if i == len(left) || j == len(right) {
			break
		}
	}
	for i < len(left) {
		newArr[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		newArr[k] = right[j]
		j++
		k++
	}
	return newArr, inversionCount
}
