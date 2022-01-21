package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ti := time.Now()
	arr := rand.Perm(100000)
	length := len(arr)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		//if min != i {
		//fmt.Println(arr[i], arr[min])
		arr[i], arr[min] = arr[min], arr[i]
		//}
	}
	//fmt.Println(arr)
	fmt.Println(time.Since(ti))
}
