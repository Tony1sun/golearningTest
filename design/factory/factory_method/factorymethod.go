package factorymethod

import "fmt"

// 工厂,New由具体工厂实现
type PenFactory interface {
	New() Pen
}

// 产品,所有具体产品的父类
type Pen interface {
	Write()
}

// 具体产品,圆珠笔
type BallPen struct{}

func (*BallPen) Write() {
	fmt.Println("BallPen Writing")
}

// 具体产品工厂
type BallPenFactory struct{}

func (*BallPenFactory) New() Pen {
	fmt.Println("create BallPen")
	return new(BallPen)
}

// 具体产品,铅笔
type Pencil struct{}

func (*Pencil) Write() {
	fmt.Println("Pencil Writing")
}

// 具体产品工厂
type PencilFactory struct{}

func (p *PencilFactory) New() Pen {
	fmt.Println("create Pencil")
	return new(Pencil)
}

func main() {
	var factory PenFactory
	factory = &PencilFactory{} // 生成一个具体产品
	var pen Pen
	pen = factory.New()
	pen.Write()

}
