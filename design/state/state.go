package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// Context 上下文
type Context struct {
	Tel string //手机号
	Text string //短信内容
	TempleteID string //短信模板ID
}

// 短信服务接口
type SmsServiceInterface interface {
	Send(ctx *Context) error
}

// 阿里云
type ServiceProviderAliyun struct {
}

// 具体发送短信的逻辑
func (s *ServiceProviderAliyun) Send(ctx *Context) error {
	fmt.Println(runFuncName(), "【阿里云】短信发送成功，手机号:"+ctx.Tel)
	return nil
}

// 腾讯云
type ServiceProviderTencent struct {
}

// 具体发送短信的逻辑
func (s *ServiceProviderTencent) Send(ctx *Context) error {
	fmt.Println(runFuncName(), "【腾讯云】短信发送成功，手机号:"+ctx.Tel)
	return nil
}

// 云片
type ServiceProvicerYunpian struct {
}

// 具体发送短信的逻辑
func (s *ServiceProvicerYunpian) Send(ctx *Context) error {
	fmt.Println(runFuncName(), "【云片】短信发送成功，手机号:"+ctx.Tel)
	return nil
}

// 获取正在运行的函数名
func runFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

// 短信服务提供商类型
type ProviderType string

const (
	// 阿里云
	ProviderTypeAliyun ProviderType = "aliyun"
	// 腾讯云
	ProviderTypeTencent ProviderType = "tencent"
	// 云片
	ProviderTypeYunpian ProviderType = "yunpian"
)

// 当前使用的服务提供商实例
// 默认aliyun
var stateManagerInstance *StateManager

// 状态管理
type StateManager struct {
	// 当前使用的服务提供商类型
	// 默认aliyun
	currentProviderTpe ProviderType

	// 当前使用的服务提供商实例
	// 默认aliyun
	currentProvider SmsServiceInterface

	// 更新状态时间间隔
	setStateDuration time.Duration
}

// 初始化状态
func (m *StateManager) initState(duration time.Duration) {
	// 初始化
	m.setStateDuration = duration
	m.setS
}

// setState 设置状态
// 根据短信云商回调的短信发送成功率 得到下阶段发送短信使用哪个厂商的服务
func (m *StateManager) setState(t time.Time) {
	// 使用随机模拟
	ProviderTypeArray := [3]ProviderType{
		ProviderTypeAliyun,
		ProviderTypeTencent,
		ProviderTypeYunpian,
	}
	m.currentProviderTpe = ProviderTypeArray[rand.Intn(len(ProviderTypeArray))]

	switch m.currentProviderTpe {
	case ProviderTypeAliyun:
		m.currentProvider = &ServiceProviderAliyun{}
	case ProviderTypeTencent:
		m.currentProvider = &ServiceProviderTencent{}
	case ProviderTypeYunpian:
		m.currentProvider = &ServiceProvicerYunpian{}
	default:
		panic("无效的短信服务商")
	}
	fmt.Printf("时间：%s| 变更短信发送厂商为: %s \n", t.Format("2006-01-02 15:04:05"), m.currentProviderType)
}