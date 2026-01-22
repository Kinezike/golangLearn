package main

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

func sum(a, b int) int {
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

func main() {
	fmt.Printf("hellow")
	nums := [5]int{1, 2, 3}

	fmt.Println(nums[0])

	length := len(nums)

	fmt.Println(length)

	//片段扩容 对齐机制
	//1.old cap * 2 < cap, newcap = cap
	//oldcap < 256 , newcap = 2 *old cap
	//
	nums1 := make([]int, 0, 0)
	nums1 = append(nums1, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(len(nums1), cap(nums1)) // 7 8 可以看到长度与容量并不一致。

	var nums2 []int
	nums2 = append(nums2, 1, 12, 3, 2, 3)
	fmt.Println(cap(nums2))

	slice := []int{1, 2, 3, 4, 5, 7, 8, 9}
	for index, val := range slice {
		fmt.Println(index, val)
	}

	mp := map[string]int{
		"a": 0,
		"b": 22,
		"c": 33,
	}
	fmt.Println(mp)
	fmt.Println(mp["a"])
	fmt.Println(mp["g"])

	newintprint := new(int)
	fmt.Println(newintprint)

	programmer := Programmer{
		Name:     "jack",
		Age:      19,
		Job:      "coder",
		Language: []string{"Go", "C++"},
	}

	fmt.Println(programmer)
	c := sum(1, 2)
	fmt.Println(c)
	fmt.Println(NamedSum(1, 2))
	fmt.Println(sum(1, 2))
}
