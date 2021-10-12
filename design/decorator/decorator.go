package decorator

import "fmt"

type Person interface {
	cost() int
	show()
}

// 被装饰对象
type laowang struct{}

func (*laowang) show() {
	fmt.Println("裸奔的老王")
}

func (*laowang) cost() int {
	return 0
}

type clothesDecorator struct {
	person Person
}

func (*clothesDecorator) cost() int {
	return 0
}

func (*clothesDecorator) show() {

}

// 夹克
type Jacket struct {
	clothesDecorator
}

func (j *Jacket) cost() int {
	return j.person.cost() + 10
}

func (j *Jacket) show() {
	// 执行已有的方法
	j.person.show()
	fmt.Println("穿上夹克的老王")
}

// 帽子
type Hat struct {
	clothesDecorator
}

func (h *Hat) cost() int {
	return h.person.cost() + 5
}

func (h *Hat) show() {
	fmt.Println("带上帽子的老王")
}
