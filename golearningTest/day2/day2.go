package day2

import (
	"fmt"
	"strconv"
)

func SortNumber(a, b int) (int, int) {
	if a > b {
		return a, b
	}
	return b, a
}

//创建一个固定大小保存整数的栈。它无须超出限制的增长。定义 push 函数—— 将数据放入栈，和 pop 函数——从栈中取得内容。栈应当是后进先出（LIFO） 的。
// 定义栈结构
type stack struct {
	node int     // 当前的节点
	data [10]int //实际数据
}

// 往栈中增加数据
func (s *stack) Push(number int) {
	if s.node < len(s.data) {
		s.data[s.node] = number
		s.node++
	}
}

// 从栈中取出数据
func (s *stack) Pop() int {
	if s.node > 0 {
		s.node--
	}
	return s.data[s.node]
}

// 更进一步。编写一个 String 方法将栈转化为字符串形式的表达。可以这样的
//方式打印整个栈： fmt.Printf("My stack %v\n", stack)
//栈可以被输出成这样的形式： [0:m] [1:l] [2:k]
func (stack *stack) ToString() string {
	var s string
	for i := 0; i < stack.node; i++ {
		k := strconv.Itoa(i)
		v := strconv.Itoa(stack.data[i])
		s += "[" + k + ":" + v + "]"
	}
	return s
}

//Tooargs 编写函数接受整数类型变参，并且每行打印一个数字。
func Tooargs(args ...int) {
	for _, v := range args {
		fmt.Println(v)
	}
}

//斐波那契数列以：1,1,2,3,5,8,13,... 开始。或者用数学形式表达：x 1 = 1;x 2 = 1;x n = x n−1 + x n−2 ∀n > 2。
// 编写一个接受 int 值的函数，并给出这个值得到的斐波那契数列。
// 从第3项开始，每一项都等于前两项之和
func Gblq(number int) []int {
	s := make([]int, number)
	for i := 1; i <= number; i++ {
		if i < 3 {
			s[i-1] = 1
		} else {
			s[i-1] = s[i-2] + s[i-3]
		}
	}
	return s
}

func Gblq1(number int) []int {
	s := make([]int, number)
	s[0], s[1] = 1, 1
	for i := 2; i < number; i++ {
		s[i] = s[i-1] + s[i-2]

	}
	return s
}

// ArrayMap 函数 map() 函数是一个接受一个函数和一个列表作为参数的函数。函
//数应用于列表中的每个元素，而一个新的包含有计算结果的列表被返回
func ArrayMap(numbers []int, f func(int) int) []int {
	for i, v := range numbers {
		numbers[i] = f(v)
	}
	return numbers
}

//MaxOf 编写一个函数，找到 int slice ( []int ) 中的最大值。
func MaxOf(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	var max = 0
	for _, v := range numbers {
		if v > max {
			max = v
		}
	}
	return max
}

//MinOf 编写一个函数，找到 int slice ( []int ) 中的最小值。
func MinOf(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	var min = numbers[0]
	for _, v := range numbers {
		if v < min {
			min = v
		}
	}
	return min
}

var pl = fmt.Println
var pf = fmt.Printf

//ArrSort 冒泡排序
func ArrSort(numbers []int) {
	len := len(numbers)
	pl("---未排序---\n", numbers)
	for i := 0; i < len; i++ {
		pl("---第", i+1, "次冒泡---")
		for j := 0; j < len-i-1; j++ {
			pl("---第", j+1, "次循环---")
			if numbers[j] > numbers[j+1] {
				numbers[j+1], numbers[j] = numbers[j], numbers[j+1]
			}
			pl(numbers)
		}
	}
	pl("---最终结果---\n", numbers)
}

//PlusTwo 编写一个函数返回另一个函数，返回的函数的作用是对一个整数 +2。函数的名称叫做 plusTwo
func PlusTwo() func(number int) int {
	v := func(number int) int {
		return number + 2
	}
	return v
}

//PlusX 使 PlusTwo 中的函数更加通用化，创建一个 plusX(y) 函数，返回一个函数用于对整数加上 x 。
func PlusX(x int) func(y int) int {
	v := func(y int) int {
		return y + x
	}
	return v
}
