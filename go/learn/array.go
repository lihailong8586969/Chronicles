package main

import "fmt"

func main() {

	// 第 1 种
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a) // 直接输出数组 [Hello World]

	// 第 2 种
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes) // 直接输出数组 [2 3 5 7 11 13]
}

