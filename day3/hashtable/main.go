package main

import (
	"fmt"
	"os"
)

//定义emp
type Emp struct {
	Id   int
	Name string
	Pre  *Emp
	Next *Emp
}

func (this *Emp) ShowMe() {
	fmt.Printf("链表%d 找到该雇员 %d \n", this.Id%7, this.Id)
}

//定义EmpLink
// 不带表头，第一个结点就存放雇员
type EmpLink struct {
	Head *Emp
}

//1. 添加员工的方法，保证添加时，编号从小到大
func (this *EmpLink) Insert(emp *Emp) {

	cur := this.Head   // 辅助指针
	var pre *Emp = nil // 辅助指针 pre 在cur前面
	// 如果当前的EmpLink就是一个空链表
	if cur == nil {
		this.Head = emp // 完成
		return
	}
	// 如果不是空链表，给emp找到对应的位置并插入
	// 让 cur和emp比较，然后让pre保持在cur前面
	for {
		if cur != nil {
			// 比较
			if cur.Id > emp.Id {
				// 找到位置
				break
			}
			pre = cur //保证同步
			cur = cur.Next
		} else {
			break
		}
	}
	// 退出时，看下是否将emp添加到链表最后
	pre.Next = emp
	emp.Next = cur
}

//显示链表的信息
func (this *EmpLink) ShowLink(no int) {
	if this.Head == nil {
		fmt.Printf("链表%d 为空\n", no)
		return
	}
	//遍历当前的链表，并显示数据
	cur := this.Head //辅助指针
	for {
		if cur != nil {
			fmt.Printf("链表%d 雇员id=%d 名字=%s", no, cur.Id, cur.Name)
			cur = cur.Next
		} else {
			break
		}
	}
	fmt.Println() // 换行
}

//定义hashtable，含有一个链表数组
type HashTable struct {
	LinkArr [7]EmpLink
}

//给HashTable 编写Insert 雇员的方法
func (this *HashTable) Insert(emp *Emp) {
	// 使用散列函数，确定将该雇员添加到哪个链表
	linkNo := this.HashFun(emp.Id)
	// 使用对应的链表添加
	this.LinkArr[linkNo].Insert(emp)
}

//编写方法，显示hashtable的所有雇员
func (this *HashTable) ShowAll() {
	for i := 0; i < len(this.LinkArr); i++ {
		this.LinkArr[i].ShowLink(i)
	}
}

//根据id查找对应的雇员,如果没有就返回nil
func (this *EmpLink) FindById(id int) *Emp {
	cur := this.Head
	for {
		if cur != nil && cur.Id == id {
			return cur
		} else if cur == nil {
			break
		}
		cur = cur.Next
	}
	return nil
}

//根据id删除对应的雇员，如果没有就返回nil
func (this *EmpLink) DeleteById(id int) *Emp {
	cur := this.Head
	flag := false

	// 如果只有一个结点
	if cur.Id == id { //只有一个结点
		this.Head = cur.Next
	}

	// 找到要删除的结点的no，和temp的下一个结点的no比较
	for {
		if cur.Next == nil { //说明到链表最后
			break
		} else if cur.Next.Id == id || cur.Id == id {
			//说明找到要删除的结点
			flag = true
			break
		}
		cur = cur.Next
	}

	if flag { //找到，删除
		cur.Next = cur.Next.Next
	} else {
		fmt.Println("sorry,要删除的id不存在")
	}
	return nil
}

//编写一个散列方法
func (this *HashTable) HashFun(id int) int {
	return id % 7 // 得到一个值，就是对应的链表的下标
}

//查找方法
func (this *HashTable) FindById(id int) *Emp {
	// 使用散列函数，确定将该雇员在哪个链表
	linkNo := this.HashFun(id)
	return this.LinkArr[linkNo].FindById(id)
}

//删除方法
func (this *HashTable) DeleteById(id int) {
	// 使用散列函数，确定将该雇员在哪个链表
	linkNo := this.HashFun(id)
	this.LinkArr[linkNo].DeleteById(id)
}

func main() {

	key := ""
	id := 0
	name := ""

	var hashtable HashTable
	for {
		fmt.Println("====雇员系统菜单====")
		fmt.Println("input 添加雇员")
		fmt.Println("show 显示雇员")
		fmt.Println("find 查找雇员")
		fmt.Println("del 删除雇员")
		fmt.Println("exit 退出系统")
		fmt.Println("请输入选项:")
		fmt.Scanln(&key)
		switch key {
		case "input":
			fmt.Println("输入雇员id")
			fmt.Scanln(&id)
			fmt.Println("输入雇员name")
			fmt.Scanln(&name)
			emp := &Emp{
				Id:   id,
				Name: name,
			}
			hashtable.Insert(emp)
		case "show":
			hashtable.ShowAll()
		case "find":
			fmt.Println("请输入id号:")
			fmt.Scanln(&id)
			emp := hashtable.FindById(id)
			if emp == nil {
				fmt.Printf("id=%d 的雇员不存在\n", id)
			} else {
				//显示雇员信息
				emp.ShowMe()
			}
		case "del":
			fmt.Println("请输入id号:")
			fmt.Scanln(&id)
			hashtable.DeleteById(id)
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("输入错误")
		}
	}
}
