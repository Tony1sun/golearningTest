package main

import (
	"fmt"
	"time"
)

// 快速排序
// 1. left表示数组左边的下标
// 2. right 表示数组左边的下标
// 3. array 表示要排序的数组
func QuickSort(left int, right int, array *[9]int) {
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
		temp = array[l]     // 把大于pivot的值 放到temp
		array[l] = array[r] // 把小于pivot的值 放到 array[l]
		array[r] = temp     // 把大于pivot的值 放到array[r]
		// array[l], array[r] = array[r], array[l]
		//优化
		if array[l] == pivot {
			r--
		}
		if array[r] == pivot {
			l++
		}
	}
	// 如果 1 == r, 再移动下
	if l == r {
		l++
		r--
	}
	// 向左递归
	if left < r {
		QuickSort(left, r, array)
	}
	// 向右递归
	if right > l {
		QuickSort(l, right, array)
	}
}

func main() {
	arr := [9]int{-9, 78, 0, 23, -567, 70, 123, 90, -23}
	fmt.Println("原始数组:", arr)
	// var arr [80000]int
	// for i := 0; i < 80000; i++ {
	// 	arr[i] = rand.Intn(900000)
	// }
	start := time.Now().Unix()
	QuickSort(0, len(arr)-1, &arr)
	end := time.Now().Unix()
	fmt.Printf("选择排序耗时=%d秒", end-start)
	// fmt.Println("main函数")
	fmt.Println("排序后数组", arr)

}
