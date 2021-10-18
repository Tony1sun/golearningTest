package main

import (
	"errors"
	"fmt"
)

//使用数组来模拟一个栈的使用
type Stack struct {
	MaxTop int    //表示栈最大可以存放的个数
	Top    int    //表示栈顶
	arr    [5]int //数组模拟栈
}

func (this *Stack) Push(val int) (err error) {
	//先判断栈是否已满
	if this.Top == this.MaxTop-1 {
		fmt.Println("stack full")
		return errors.New("stack full")
	}
	this.Top++
	//放入数据
	this.arr[this.Top] = val
	return
}

// 遍历栈，需要从栈顶开始遍历
func (this *Stack) List() {
	// 先判断栈是否为空
	if this.Top == -1 {
		fmt.Println("stack empty")
		return
	}
	fmt.Println("栈的情况如下:")
	for i := this.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, this.arr[i])
	}
}

func main() {
	stack := &Stack{
		MaxTop: 5,  //表示最多存放5个数到栈中
		Top:    -1, //当栈顶为-1，表示栈为空
	}
	//入栈
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)
	//显示
	stack.List()
}