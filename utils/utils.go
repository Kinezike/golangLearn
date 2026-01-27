package utils

import (
	"bufio"
	"errors"
	"fmt"
	"hellowgo/example"
	"hellowgo/structlearning"
	"math"
	"os"
)

type Programmer struct {
	Name     string
	Age      int
	Job      string
	Language []string
}

func Sum2(a, b int) int {
	return a + b
}

// 2. 多返回值 - 无名返回值
func Div(a, b float64) (float64, error) {
	if b == 0 {
		return math.NaN(), errors.New("0不能作为除数")
	}
	return a / b, nil
}

func NamedSum(a, b int) (ans int) {
	return a + b
}

func printlearn() {
	res := example.SayHello("Go") // 通过包名调用函数
	fmt.Println(res)
}

func bufiolearn() {
	/* 缓冲 */

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println(scanner.Text())
}

func structlearn() {
	/* struct学习 */
	v1 := structlearning.Person{
		Name: "mike",
		Age:  10,
	}
	fmt.Printf("%+v", v1)
}

func inputlearn() {
	/* 输入 */
	var s1, s2 string
	fmt.Scan(&s1, &s2)
	fmt.Print(s1, s2)
}

func ifelselearn() {
	a, b := 2, 2
	if a < b {
		fmt.Print("a < b")
	} else if a == b {
		fmt.Print("a = b")
	} else {
		fmt.Print("a > b")
	}

	str := "a"
	switch str {
	case "a":
		str += "a"
		str += "c"
	case "b":
		str += "bb"
		str += "aaaa"
	default: // 当所有case都不匹配后，就会执行default分支
		str += "CCCC"
	}
	fmt.Println(str)
}

func forlearn() {
	/* for i := 0; i <= 20; i++ {
		fmt.Println(i)
	} */

	sequence := "hello world"
	for index, value := range sequence {
		fmt.Println(index, value)
	}
}

/****************************************************************/
func Sum(a, b int) int {
	return a + b
}

func Sum1(a, b int) (ans int) {
	ans = a + b
	return
}

/****************************************************************/

// 命名返回值：sum 和 avg 是返回值别名
func calcSumAvg(nums []int) (sum int, avg float64) {
	total := 0
	for _, num := range nums {
		total += num
	}
	sum = total // 直接为命名返回值赋值
	avg = float64(total) / float64(len(nums))
	return // 隐式返回 sum 和 avg（无需写 return sum, avg）
}

func process() (result int) { // result 是命名返回值（别名）
	defer func() {
		if result < 100 {
			result = 100 // defer 中修改命名返回值
		}
	}()
	//可以直接使用defer在函数体当中去修改返回值

	result = 50 // 初始赋值
	return      // 最终返回被 defer 修改为 100
}

func (p Programmer) Sayhi(alice string) (ans int) {
	ans = 1
	fmt.Println("hello,i`m", p.Name, alice)
	return
}
