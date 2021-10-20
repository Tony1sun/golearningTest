package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 选择排序
// 编写函数selectSort 完成排序
// 注释
func SelectSort(arr *[800000]int) {
	// 1.假设 arr[0] 是最大值
	for j := 0; j < len(arr)-1; j++ {
		max := arr[j]
		maxIndex := j
		// 2.遍历后面 1----[len(arr) - 1] 比较
		for i := j + 1; i < len(arr); i++ {
			if max < arr[i] {
				max = arr[i]
				maxIndex = i
			}
		}
		// 交换
		// 如果最大值的下标不等于j
		if maxIndex != j {
			arr[j], arr[maxIndex] = arr[maxIndex], arr[j]
		}
		// fmt.Println(arr)
		fmt.Println(arr)
	}
}

func main() {
	// 定义一个数组
	// arr := [6]int{10, 34, 19, 100, 80, 789}
	var arr [800000]int
	for i := 0; i < 800000; i++ {
		arr[i] = rand.Intn(900000)
	}
	start := time.Now().Unix()
	SelectSort(&arr)
	end := time.Now().Unix()
	fmt.Printf("选择排序耗时=%d秒", end-start)
	// fmt.Println("main函数")
	// fmt.Println(arr)
}
