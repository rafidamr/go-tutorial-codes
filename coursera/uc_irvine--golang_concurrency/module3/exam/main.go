package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {
	var arrLen int
	fmt.Scanf("%d", &arrLen)

	arr := make([]int, arrLen)
	for i := range arr {
		fmt.Scanf("%d", &arr[i])
	}

	subArrays := splitArray(arr)
	concurrentSort(subArrays)
	arrX := mergeTwoSubArrays(subArrays[0], subArrays[1])
	arrY := mergeTwoSubArrays(subArrays[2], subArrays[3])
	arrFinal := mergeTwoSubArrays(arrX, arrY)
	fmt.Printf("Final Array: %v\n", arrFinal)
}

func splitArray(arr []int) [4][]int {
	n := len(arr)

	subArrSize := n / 4
	var subArrays [4][]int

	for i := 0; i < 4; i++ {
		start := i * subArrSize
		end := start + subArrSize

		if i == 3 {
			end = n
		}

		subArrays[i] = make([]int, end-start)
		copy(subArrays[i], arr[start:end])
	}

	return subArrays
}

func concurrentSort(subArrays [4][]int) {
	var wg sync.WaitGroup
	wg.Add(4)
	for i := range subArrays {
		go func(idx int) {
			defer wg.Done()
			fmt.Printf("subArray %v: %v\n", i, subArrays[i])
			sort.Ints(subArrays[i])
		}(i)
	}
	wg.Wait()
}

func mergeTwoSubArrays(arr1 []int, arr2 []int) []int {
	combinedSize := len(arr1) + len(arr2)
	result := make([]int, 0, combinedSize)
	i, j := 0, 0

	for idx := 0; idx < combinedSize; idx++ {
		if i < len(arr1) && j < len(arr2) {
			if arr1[i] <= arr2[j] {
				result = append(result, arr1[i])
				i++
			} else {
				result = append(result, arr2[j])
				j++
			}
		}
	}

	result = append(result, arr1[i:]...)
	result = append(result, arr2[j:]...)

	return result
}
