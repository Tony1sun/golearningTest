package main

import "fmt"

func InsertSort(arr *[7]int) {

	for i := 1; i < len(arr); i++ {
		// 完成第一次，给第二个元素找到合适的位置，并插入
		insertVal := arr[i]
		insertIndex := i - 1

		// 从大到小
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex] // 后移
			insertIndex--
		}
		//插入
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertVal

		}
		fmt.Printf("第%d次插入后 %v\n", i, *arr)
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
	arr := [7]int{23, 0, 12, 56, 34, -1, 55}
	fmt.Println("原始数组是:", arr)
	InsertSort(&arr)
	fmt.Println("main函数")
	fmt.Println(arr)
}
