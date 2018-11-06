package main

import "fmt"

func main() {

	primes := [6]int{2, 3, 5, 7, 11, 13}

	// 切片并不存储任何数据，它只是描述了底层数组中的一段。
	// 更改切片的元素会修改其底层数组中对应的元素。
	// 与它共享底层数组的切片都会观测到这些修改。
	// a[low : high] 左闭右开
	var s []int = primes[1:4]
	fmt.Println(s)


	s := []int{2, 3, 5, 7, 11, 13}

	// from index 1 to 4-1
	s = s[1:4]
	fmt.Println(s)

	// from index 0 to 2-1
	s = s[:2]
	fmt.Println(s)

	// 从下标1开始 to end
	s = s[1:]
	fmt.Println(s)

}
