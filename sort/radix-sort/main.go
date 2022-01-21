package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	arr := rand.Perm(100)
	fmt.Println(RadixSort(arr))
}

func RadixSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	maxNum := nums[0] //最大数
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}

	// maxDigit := 0
	// for maxNum > 0 {
	// 	maxNum = maxNum / 10
	// 	maxDigit++
	// }
	// fmt.Println(maxDigit)

	maxDight := int(math.Log10(float64(maxNum))) + 1 // 最大位数

	fmt.Println(maxDight)

	for i, j := 1, 1; i <= maxDight; i++ {
		bucket := [10][]int{}
		for _, v := range nums {
			k := (v / j) % 10
			bucket[k] = append(bucket[k], v)
		}
		j *= 10

		k := 0
		for _, v := range bucket {
			for _, v1 := range v {
				nums[k] = v1
				k++
			}
		}
	}

	return nums
}
