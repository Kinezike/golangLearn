package test

import (
	"fmt"
	"hellowgo/tree"
)

type Person interface {
	Say(string) string
	Walk(int)
}

type Logger interface {
	Log(string)
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

func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	ch <- t.Value
	Walk(t.Left, ch)
	Walk(t.Right, ch)
	n := cap(ch)
	for i := 1; i < n; i++ {
		fmt.Println(ch)
	}
	close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()

	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if ok1 != ok2 {
			return false
		}
		if !ok1 {
			return true
		}
		if v1 != v2 {
			return false
		}

	}
}
func TEst() {
	ch1 := make(chan int)

	go func() {
		for i := 1; i < 5; i++ {
			ch1 <- i
		}
	}()

	go func() {
		for i := 1; i < 5; i++ {
			ch1 <- 8
		}
	}()

	for i := range ch1 {
		println("ch1:=%d", i)
	}
}
