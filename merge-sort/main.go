package main

import "fmt"

// 归并排序
func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2 // 取中间位置
	left := mergeSort(arr[:mid]) // 左边的数据进行排序
	right := mergeSort(arr[mid:]) // 右边的数据进行排序
	return merge(left, right) // 左右两边的数据进行合并
}

func merge(left, right []int) []int {
	result := make([]int, 0)
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] { // 左边的数比右边的数小
			result = append(result, left[0]) // 左边的数放入结果
			left = left[1:]                  // 左边的数继续取
		} else {
			result = append(result, right[0]) // 右边的数放入结果
			right = right[1:]                 // 右边的数继续取
		}
	}
	result = append(result, left...)  // 左边的数放入结果
	result = append(result, right...) // 右边的数放入结果
	return result                     // 返回结果
}

func main() {
	arr := []int{5, 4, 3, 2, 1, 432, 13, 4, 51}
	re := mergeSort(arr)
	fmt.Println(re)
}
