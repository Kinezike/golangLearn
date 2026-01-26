package main

import "fmt"

// 外部函数：接收一个参数 base（环境的一部分）
func createAdder(base int) func(int) int {
	// 内部函数（闭包）：捕获外部参数 base，并与自身参数 x 相加
	return func(x int) int {
		return base + x // 使用捕获的 base 和自身参数 x
	}
}

func main() {
	// 构造两个不同的闭包实例（分别捕获 base=5 和 base=10）
	add5 := createAdder(5)   // 闭包1：捕获 base=5
	add10 := createAdder(10) // 闭包2：捕获 base=10

	// 调用闭包：每个闭包使用自己捕获的 base
	fmt.Println(add5(3))  // 5+3=8
	fmt.Println(add5(7))  // 5+7=12
	fmt.Println(add10(3)) // 10+3=13
}
