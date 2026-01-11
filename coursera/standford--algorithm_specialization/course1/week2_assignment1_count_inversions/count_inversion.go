package main

import "fmt"

func main() {
	arr := []int{1, 100, 2, 10}
	sortedArr := mergeSort(arr)
	fmt.Println(sortedArr)
}

func mergeSort(arr []int) []int {
	_len := len(arr)
	if _len == 1 {
		return arr
	}
	left := mergeSort(arr[:(_len / 2)])
	right := mergeSort(arr[(_len / 2):])
	newArr := merge(left, right)
	return newArr
}

func merge(left []int, right []int) []int {
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
	return newArr
}
