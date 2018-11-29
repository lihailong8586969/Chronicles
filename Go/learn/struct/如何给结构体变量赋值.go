package main

import "fmt"

type Man struct {

	name string
	age int8
}

func main() {

	// {"name":"", "age":32} [ X ] 因为不能加引号
	man  := &Man{name:"Eason", age:33} // 开辟空间初始化并返回这个空间的指针
	man2 := Man{name:"jay", age:23} // 开辟空间且初始化并返回这个变量
	var man3 Man // 定义结构体变量
	var man4 Man // 定义结构体变量
	man4 = Man{} // 给结构体变量赋值

	fmt.Println(man)  // man  是指针类型
	fmt.Println(man2) // man2 是变量类型
	fmt.Println(man3) // man3 是变量类型


}