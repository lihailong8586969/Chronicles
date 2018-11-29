package main

import "fmt"

type Man struct {

	name string
	age int8
}

func main() {

	// man 是一个指针，和 C 语言中的一样
	man := new(Man)

	fmt.Println(man)
	fmt.Println("age : " , (*(man)).age)
}