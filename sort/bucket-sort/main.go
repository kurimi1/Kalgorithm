package main

import "fmt"

// bucket sort

func BucketSort(arr []int) []int {
	// 定义桶数
	num := len(arr)

	// 找到k（数组最大值）
	max := getMaxInArr(arr)

	// 二维切片 (维度1表示桶的编号,维度2是一个数组,存放对应桶的元素)
	buckets := make([][]int, num)

	//分配入桶
	index := 0
	for i := 0; i < num; i++ {
		index = arr[i] * (num - 1) / max // 将待排序的元素分配桶  index = value * (n-1) /k
		buckets[index] = append(buckets[index], arr[i])
	}

	//桶内排序
	tmpPos := 0
	for i := 0; i < num; i++ {
		bucketLen := len(buckets[i])
		if bucketLen > 0 {
			InsertSort(buckets[i])
			copy(arr[tmpPos:], buckets[i])
			tmpPos += bucketLen
		}
	}

	return arr
}

//获取数组最大值
func getMaxInArr(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

// InsertSort 插入排序算法
func InsertSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	for i := 1; i < length; i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}
	return arr
}

func main() {
	arr := []int{1, 3, 23, 45, 14, 55, 55, 343, 2, 2, 5, 4, 6, 7, 8, 9, 10}
	BucketSort(arr)
	fmt.Println(arr)
}
