package test

import "fmt"

type Person interface {
	Say(string) string
	Walk(int)
}

// 外部函数：接收一个参数 base（环境的一部分）
func createAdder(base int) func(int) int {
	// 内部函数（闭包）：捕获外部参数 base，并与自身参数 x 相加
	return func(x int) int {
		return base + x // 使用捕获的 base 和自身参数 x
	}
}

func Adder(base int) func(int) int {
	return func(x int) int {
		return base + x
	}
}

func Do() {
	defer func() {
		fmt.Println("1")
	}()
	fmt.Println("2")
}
