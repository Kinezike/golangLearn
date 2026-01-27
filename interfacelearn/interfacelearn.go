package interfacelearn

import "fmt"

// 起重机接口
type Crane interface {
	JackUp() string
	Hoist() string
}

// 起重机A
type CraneA struct {
	work int //内部的字段不同代表内部细节不一样
}

func (c CraneA) Work() {
	fmt.Println("使用技术A")
}
func (c CraneA) JackUp() string {
	c.Work()
	return "jackup"
}

func (c CraneA) Hoist() string {
	c.Work()
	return "hoist"
}

// 起重机B
type CraneB struct {
	boot string
}

func (c CraneB) Boot() {
	fmt.Println("使用技术B")
}

func (c CraneB) JackUp() string {
	c.Boot()
	return "jackup"
}

func (c CraneB) Hoist() string {
	c.Boot()
	return "hoist"
}

type ConstructionCompany struct {
	Crane Crane // 只根据Crane类型来存放起重机
}

func (c *ConstructionCompany) Build() {
	fmt.Println(c.Crane.JackUp())
	fmt.Println(c.Crane.Hoist())
	fmt.Println("建筑完成")
}
