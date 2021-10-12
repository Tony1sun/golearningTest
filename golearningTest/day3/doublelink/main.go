package main

import (
	"fmt"
)

// 定义一个HeroNode
type HeroNode struct {
	no       int
	name     string
	nickname string
	pre      *HeroNode // 指向前一个结点
	next     *HeroNode // 指向下一个结点
}

//给链表插入一个结点
//在单链表最后加入-第一种方法
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	//1. 先找到该链表的最后结点
	//2. 创建一个辅助结点
	temp := head
	for {
		if temp.next == nil { // 找到最后
			break
		}
		temp = temp.next //让temp不断的指向下一个结点
	}
	//3. 将newHeroNode加入到链表的最后
	temp.next = newHeroNode
	newHeroNode.pre = temp
}

//给双向链表插入一个结点
//根据no的编号从小到大插入...
func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	//1. 找到适当的结点
	//2. 创建一个辅助结点
	temp := head
	flag := true
	// 让插入的结点的no，和temp的下一个结点的no比较
	for {
		if temp.next == nil { // 说明到链表的最后
			break
		} else if temp.next.no > newHeroNode.no {
			// 说明newHeroNode 应该插入到temp后面
			break
		} else if temp.next.no == newHeroNode.no {
			// 说明链表中已经存在这个no
			flag = false
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Println("已经存在no=", newHeroNode.no)
		return
	} else {
		newHeroNode.next = temp.next
		newHeroNode.pre = temp
		if temp.next != nil {
			temp.next.pre = newHeroNode
		}
		temp.next = newHeroNode

	}

}

// 删除一个结点
func DelHerNode(head *HeroNode, id int) {
	temp := head
	flag := false

	// 找到要删除的结点的no，和temp的下一个结点的no比较
	for {
		if temp.next == nil { // 说明到链表最后
			break
		} else if temp.next.no == id {
			// 说明找到要删除的结点
			flag = true
			break
		}
		temp = temp.next
	}

	if flag { // 找到，删除
		temp.next = temp.next.next
		if temp.next != nil {
			temp.next.pre = temp
		}
	} else {

		fmt.Println("sorry, 要删除的id不存在!")

		fmt.Println("sorry,要删除的id不存在")
	}

}

// 修改一个结点
func modifyStudentNode(head *HeroNode, id int, newHeroNode *HeroNode) {
	//1. 找到适当的结点
	//2. 创建一个辅助结点
	temp := head
	flag := false
	// 找到要修改的结点的no,和temp.next.no作比较
	for {
		if temp.next == nil { // 说明到链表的最后
			break
		} else if temp.next.no == id {
			// 说明已经找到要修改的结点
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		temp.next = newHeroNode
	} else {
		fmt.Println("要修改的结点不存在!")

	}
}

// 显示链表所有结点信息
func ListHeroNode(head *HeroNode) {
	//1. 创建一个辅助结点
	temp := head

	// 先判断该链表是不是一个空的链表
	if temp.next == nil {
		fmt.Println("这是一个空链表......")
		return
	}

	//2. 遍历这个链表
	for {
		fmt.Printf("[%d, %s, %s]==>", temp.next.no,
			temp.next.name, temp.next.nickname)
		// 判断是否到链表最后
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}

func ListHeroNode2(head *HeroNode) {
	//1. 创建一个辅助结点
	temp := head

	// 先判断该链表是不是一个空的链表
	if temp.next == nil {
		fmt.Println("这是一个空链表......")
		return
	}

	//2. 让temp定位到双向链表的最后结点
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}

	//2. 遍历这个链表
	for {
		fmt.Printf("[%d, %s, %s]==>", temp.no,
			temp.name, temp.nickname)
		// 判断是否到链表头
		temp = temp.pre
		if temp.pre == nil {
			break
		}
	}
}

func main() {

	//1. 先创建一个头结点
	head := &HeroNode{}

	//2. 创建一个新的HeroNode
	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
	}
	hero2 := &HeroNode{
		no:       2,
		name:     "鲁智深",
		nickname: "花和尚",
	}
	hero3 := &HeroNode{
		no:       3,
		name:     "郭敬明",
		nickname: "迅捷斥候",
	}

	//3. 加入

	InsertHeroNode2(head, hero1)
	InsertHeroNode2(head, hero2)
	InsertHeroNode2(head, hero3)
	//4. 显示
	ListHeroNode(head)
	fmt.Println("逆序打印")
	ListHeroNode2(head)

	// fmt.Println()
	// //5. 删除
	// DelHerNode(head, 1)
	// ListHeroNode(head)
	stuMarket := &HeroNode{
		no:       3,
		name:     "卡牌大师",
		nickname: "刘谦",
	}

	//3. 加入
	InsertHeroNode(head, hero3)

	InsertHeroNode(head, hero1)
	InsertHeroNode(head, hero2)

	//4. 显示
	// ListHeroNode(head)
	// fmt.Println("逆序打印")
	ListHeroNode(head)
	//删除结点
	DelHerNode(head, 1)
	fmt.Println()
	ListHeroNode(head)
	fmt.Println()

	//修改结点
	modifyStudentNode(head, 2, stuMarket)
	ListHeroNode(head)

}
