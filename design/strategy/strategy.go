package strategy

import (
	"fmt"
	"runtime"
)

// 策略模式
const (
	ConstWechatPay = "wechat_pay"
	ConstAliPayWap = "AliPayWapwap"
	ConstBankPay   = "quickbank"
)

// Context上下文
type Context struct {
	// 用户选择的输入方式
	PayType string `json:"pay_type"`
}

type PaymentInterface interface {
	Pay(ctx *Context) error    //支付
	Refund(ctx *Context) error //退款
}

// 微信支付
type WechatPay struct{}

// 当前支付方式的支付逻辑
func (p *WechatPay) Pay(ctx *Context) (err error) {
	// 当前策略的业务逻辑
	fmt.Println(runFuncName(), "使用微信支付...")
	return
}

func (p *WechatPay) Refund(ctx *Context) (err error) {
	// 当前策略的业务逻辑
	fmt.Println(runFuncName(), "使用微信退款...")
	return
}

// 支付宝网页版
type AliPayWap struct{}

func (p *AliPayWap) Pay(ctx *Context) (err error) {
	// 当前策略的业务逻辑
	fmt.Println(runFuncName(), "使用支付宝网页版支付...")
	return
}

func (p *AliPayWap) Refund(ctx *Context) (err error) {
	// 当前策略的业务逻辑
	fmt.Println(runFuncName(), "使用支付宝退款...")
	return
}

// BankPay 银行卡支付
type BankPay struct {
}

// Pay 当前支付方式的支付逻辑
func (p *BankPay) Pay(ctx *Context) (err error) {
	// 当前策略的业务逻辑写这
	fmt.Println(runFuncName(), "使用银行卡支付...")
	return
}

// Refund 当前支付方式的支付逻辑
func (p *BankPay) Refund(ctx *Context) (err error) {
	// 当前策略的业务逻辑写这
	fmt.Println(runFuncName(), "使用银行卡退款...")
	return
}

// 获取正在运行的函数名
func runFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

func main() {
	// 业务上下文
	ctx := &Context{
		PayType: "wechat_pay",
	}

	// 获取支付方式
	var instance PaymentInterface
	switch ctx.PayType {
	case ConstWechatPay:
		instance = &WechatPay{}
	case ConstAliPayWap:
		instance = &AliPayWap{}
	case ConstBankPay:
		instance = &BankPay{}
	default:
		panic("无效的支付方式")
	}
	//支付
	instance.Pay(ctx)
}
