package main

import "fmt"

// func InsertSort(arr []int) {

// 	for i := 1; i < len(arr); i++ {
// 		// 完成第一次，给第二个元素找到合适的位置，并插入
// 		insertVal := arr[i]
// 		insertIndex := i - 1

// 		// 每次取到的数都跟左侧的数做比较，如果左侧的数比取的数大，就将左侧的数右移一位，直至左侧没有数字比取的数大为止
// 		for insertIndex >= 0 && arr[insertIndex] < insertVal {
// 			arr[insertIndex+1] = arr[insertIndex] //数据后移
// 			insertIndex--                         //往前移动
// 		}
// 		//插入
// 		if insertIndex+1 != i {
// 			arr[insertIndex+1] = insertVal

// 		}
// 		fmt.Printf("第%d次插入后 %v\n", i, arr)
// 	}

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
// }

func main() {
	arr := []int{23, 0, 56, 34, -1, 55}
	fmt.Println("原始数组是:", arr)
	InsertSort(arr)
	fmt.Println("main函数")
	fmt.Println(arr)
}

func InsertSort(arr []int) {
	// 完成第一次，给第二个值找插入位置
	for i := 1; i < len(arr); i++ {
		insertVal := arr[i]
		insertIndex := i - 1
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}
		// 交换
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertVal
		}

		fmt.Printf("第%d次排序后的值%v\n", i, arr)
	}
}
