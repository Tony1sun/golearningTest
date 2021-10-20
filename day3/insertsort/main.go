package main

import (
	"fmt"
	"math/rand"
	"time"
)

<<<<<<< HEAD
func InsertSort(arr []int) {
=======
func InsertSort(arr *[80000]int) {
>>>>>>> b2f0b380627410b235c225410b9c29d9e80a17a6

	for i := 1; i < len(arr); i++ {
		// 完成第一次，给第二个元素找到合适的位置，并插入
		insertVal := arr[i]
		insertIndex := i - 1

		// 从大到小
		for insertIndex >= 0 && arr[insertIndex] < insertVal { //insertVal大于insertIndex，那insertIndex就往前移一位
			arr[insertIndex+1] = arr[insertIndex] //数据后移
			insertIndex--                         //往前移动
		}
		//插入
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertVal

		}
<<<<<<< HEAD
		fmt.Printf("第%d次插入后 %v\n", i, arr)
=======
		// fmt.Printf("第%d次插入后 %v\n", i, *arr)
>>>>>>> b2f0b380627410b235c225410b9c29d9e80a17a6
	}

	// // 完成第2次，给第3个元素找到合适的位置，并插入
	// insertVal = arr[2]
	// insertIndex = 2 - 1

	// // 从大到小
	// for insertIndex >= 0 && arr[insertIndex] < insertVal {
	// 	arr[insertIndex+1] = arr[insertIndex] // 后移
	// 	insertIndex--
	// }
	// //插入
	// if insertIndex+1 != 2 {
	// 	arr[insertIndex+1] = insertVal

	// }
	// fmt.Println("第2次插入后", *arr)

	// // 完成第3次，给第4个元素找到合适的位置，并插入
	// insertVal = arr[3]
	// insertIndex = 3 - 1

	// // 从大到小
	// for insertIndex >= 0 && arr[insertIndex] < insertVal {
	// 	arr[insertIndex+1] = arr[insertIndex] // 后移
	// 	insertIndex--
	// }
	// //插入
	// if insertIndex+1 != 3 {
	// 	arr[insertIndex+1] = insertVal
	// }
	// fmt.Println("第3次插入后", *arr)
}

func main() {
<<<<<<< HEAD
	arr := []int{23, 0, 12, 56, 34, -1, 55}
	fmt.Println("原始数组是:", arr)
	InsertSort(arr)
	fmt.Println("main函数")
	fmt.Println(arr)
=======
	// arr := [80000]int{23, 0, 12, 56, 34, -1, 55}
	// fmt.Println("原始数组是:", arr)
	var arr [80000]int
	for i := 0; i < 80000; i++ {
		arr[i] = rand.Intn(900000)
	}
	start := time.Now().Unix()

	InsertSort(&arr)
	end := time.Now().Unix()
	fmt.Printf("选择排序耗时=%d秒", end-start)

	// fmt.Println("main函数")
	// fmt.Println(arr)
>>>>>>> b2f0b380627410b235c225410b9c29d9e80a17a6
}
