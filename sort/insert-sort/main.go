package main

import (
	"fmt"
)

func InsertSort(arr []int) []int {
	for i := range arr {
		preIndex := i - 1 // 1
		fmt.Println(preIndex)
		current := arr[i]                              // 与前一个相比                              // a[2]
		for preIndex >= 0 && arr[preIndex] > current { // a[1] > a[2]
			arr[preIndex+1] = arr[preIndex] // a[2] = a[1]
			preIndex--
		}
		arr[preIndex+1] = current // a[1] = a[2]
	}
	return arr
}

func main() {
	arr := []int{5, 1, 3}
	fmt.Println(InsertSort(arr))
}
