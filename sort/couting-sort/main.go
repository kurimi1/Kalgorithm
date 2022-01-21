package main

import "fmt"

// 计数排序
// 时间复杂度：O(n+k)
// 空间复杂度：O(n+k)
func countingSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	// 计算最大值
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	// 计算每个值出现的次数
	count := make([]int, max+1)
	for _, v := range arr {
		count[v]++
	}

	// 将计数结果取出
	index := 0
	for i := 1; i < len(count); i++ {
		for count[i] > 0 {
			arr[index] = i
			index++
			count[i]--
		}
	}

	return arr
}

func main() {
	arr := []int{1, 3, 23, 45, 14, 55, 2, 2, 5, 4, 6, 7, 8, 9, 10}
	result := countingSort(arr)
	fmt.Println(result)
}
