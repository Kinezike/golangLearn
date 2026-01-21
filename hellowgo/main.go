package main

import "fmt"

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
}
