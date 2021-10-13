package main

import "fmt"

// 快速排序
// 1. left表示数组左边的下标
// 2. right 表示数组左边的下标
// 3. array 表示要排序的数组
func QuickSort(left int, right int, array *[6]int) {
	l := left
	r := right
	// pivot 是中轴，支点
	pivot := array[(left+right)/2]
	temp := 0
	// 将比pivot小的数放到左边
	// 比pivot大的数放到右边
	for l < r {
		// 先从pivot的左边找到大于等于pivot的值
		for array[l] < pivot {
			l++
		}
		// 先从pivot的右边找到小于等于pivot的值
		for array[r] > pivot {
			r--
		}
		// l >= r 表明本次分解任务完成
		if l >= r {
			break
		}
		// 交换
		temp = array[l] // 把 大于pivot的值 放到temp
		array[l] = array[r]
		array[r] = temp
		//优化
		if array[l] == pivot {
			r--
		}
		if array[r] == pivot {
			l++
		}
	}

}

func main() {
	arr := [6]int{-9, 78, 0, 23, -567, 70}
	QuickSort(0, len(arr)-1, &arr)
	fmt.Println("main函数")
	fmt.Println(arr)

}
