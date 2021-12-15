package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 2, 4, 6, 1, 3}
	ShellSort(arr)
	fmt.Println(arr)
	arr2 := []int{5, 2, 4, 6, 1, 3}
	ShellSort2(arr2)
	fmt.Println(arr2)
}

func ShellSort(arr []int) []int {
	length := len(arr)
	gap := 1
	for gap < length/3 { // gap = 1, 4, 13, 40, 121, 364, 1093, ...
		gap = gap*3 + 1
	}
	for gap > 0 {
		for i := gap; i < length; i++ { // 1 , 3
			temp := arr[i]                // a[1]
			j := i - gap                  // 0
			for j >= 0 && arr[j] > temp { // a[0] > a[1]
				arr[j+gap] = arr[j] // a[1] = a[0]
				j -= gap
			}
			arr[j+gap] = temp // a[0] = a[1]
		}
		gap = gap / 3
	}
	return arr
}

// 动态分组
func ShellSort2(arr []int) []int {
	length := len(arr)
	for gap := length / 2; gap > 0; gap = gap / 2 {
		// fmt.Println(gap)
		for i := gap; i < length; i++ {
			j := i // 3
			fmt.Println(j, "1")
			temp := arr[i]
			for j-gap >= 0 && temp < arr[j-gap] {
				arr[j] = arr[j-gap]
				j = j - gap
				fmt.Println(j)
			}
			arr[j] = temp
		}
	}
	return arr
}
