package main

import "fmt"

type Man struct{

	name string;
	age int;
}

type Vertex struct {
	X int
	Y int
}


func main() {

	// 初始化字段不一定要全部指定，比如下面也是可以的，name默认取长度为0的空字符串
	man := Man{"jayChou"}

	// 全部字段都初始化
	man := Man{name:"jayChou",age:20}
	fmt.Println(man.name)

	// var man Man 声明
	var man Man = Man{"jayCHou",30}

	v := Vertex{1, 2}	
	v.X = 4
	p := &v
	p.X = 1e9	
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex



	fmt.Println("Hello, 世界")
}
