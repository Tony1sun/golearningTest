package abstract

import "fmt"

// 抽象工厂模式
type Girl interface {
	weight()
}

// 中国胖女孩
type FatGirl struct{}

func (FatGirl) weight() {
	fmt.Println("chinese girl weight: 80kg")
}

// 瘦女孩
type ThinGirl struct{}

func (ThinGirl) weight() {
	fmt.Println("chinese girl weight: 50kg")
}

type Factory interface {
	CreateGirl(like string) Girl
}

// 中国工厂
type ChineseGirlFactory struct{}

func (ChineseGirlFactory) CreateGirl(like string) Girl {
	if like == "fat" {
		return &FatGirl{}
	} else if like == "thin" {
		return &ThinGirl{}
	}
	return nil
}

// 美国胖女孩
type AmericanFatGirl struct{}

func (AmericanFatGirl) weight() {
	fmt.Println("American girl weight: 80kg")

}

// 美国瘦女孩
type AmericanThinGirl struct{}

func (AmericanThinGirl) weight() {
	fmt.Println("American girl weight: 50kg")
}

// 美国工厂
type AmericangirlFactory struct{}

func (AmericangirlFactory) CreateGirl(like string) Girl {
	if like == "fat" {
		return &AmericanFatGirl{}
	} else if like == "thin" {
		return &AmericanThinGirl{}
	}
	return nil
}

// 工厂提供者
type GirlFactoryStore struct {
	factory Factory
}

func (store *GirlFactoryStore) createGirl(like string) Girl {
	return store.factory.CreateGirl(like)
}
