package main

import "fmt"

// 小孩的结构体
type Boy struct {
	No   int  //编号
	Next *Boy //指向下一个小孩的指针
}

//编写一个函数，构成单向的环形链表
//num 小孩的个数
//*Boy 返回该环形链表第一个小孩的指针
func AddBoy(num int) *Boy {

	first := &Boy{}  // 空结点
	curBoy := &Boy{} // 空结点

	if num < 1 {
		fmt.Println("num值小于1")
		return first
	}
	//循环构建循环链表
	for i := 1; i <= num; i++ {
		boy := &Boy{
			No: i,
		}
		// 构成循环链表，需要一个辅助指针
		// 因为第一个小孩比较特殊
		if i == 1 {
			first = boy
			curBoy = boy
			curBoy.Next = first
		} else {
			curBoy.Next = boy
			curBoy = boy
			curBoy.Next = first // 构造环形链表
		}
	}
	return first
}

//显示单向的环形链表
func ShowBoy(first *Boy) {

	// 如果环形链表为空
	if first == nil {
		fmt.Println("空链表无法遍历!")
		return
	}

	// 创建一个指针，帮助遍历
	curBoy := first
	for {
		fmt.Printf("小孩编号=%d->", curBoy.No)
		//退出的条件
		if curBoy.Next == first {
			// 显示到链表最后的小孩
			break
		}
		// curBoy移动到下一个
		curBoy = curBoy.Next
	}
}

/* 设编号1，2....n的n个人围坐一圈，约定编号为k(1<=k<=n)
的人从1开始报数，数到m的那个人出列，它的下一位又开始从1开始报数,
数到m的那个人又出列，依次类推，直到所有人出列位置，由此产生一个出队编号的序列
*/

//1. 编写一个函数，PlayGame(first *Boy, startNo int, countNum int)
//2. 在环形链表中留下最后一个人
func PlayGame(first *Boy, startNo int, countNum int) {

	// 1.空链表单独处理
	if first.Next == nil {
		fmt.Println("就一个小孩玩游戏(空链表),玩不了!")
		return
	}
	// 判断 startNo <= 小孩的总数
	// 2.定义需要辅助的指针，帮助我们删除小孩
	tail := first
	// 3.让tail指向环形链表的最后一个小孩
	// 因为tail 在删除小孩时需要使用到
	for {
		if tail.Next == first { // 说明tail到了最后的小孩
			break
		}
		tail = tail.Next
	}
	//4. 让first 移动到 startNo【后面删除小孩，就以first为准】
	for i := 1; i <= startNo-1; i++ {
		first = first.Next
		tail = tail.Next
	}

	fmt.Println()
	//5. 开始数countNum，然后就删除first指向的小孩
	for {
		// 开始数countNum-1次
		for i := 1; i <= countNum-1; i++ {
			first = first.Next
			tail = tail.Next
		}
		fmt.Printf("小孩编号为%d 出圈 \n", first.No)
		// 删除first指向的小孩
		first = first.Next
		tail.Next = first
		// 判断如果 tail = first,圈中只有一个小孩
		if tail == first {
			break
		}
	}
	fmt.Printf("最后出圈的小孩编号为%d \n", first.No)
}

func main() {

	first := AddBoy(5)
	// 显示
	ShowBoy(first)
	PlayGame(first, 2, 3)
}