package main

import "fmt"

type Man struct{

	name string;
	age int;
}

// Go 没有类。不过你可以为结构体类型定义方法。
// 方法就是一类带特殊的 接收者 参数的函数。
// 方法接收者在它自己的参数列表内，位于 func 关键字和方法名之间。
// 方法只是个带接收者参数的函数。

// 对数据的修改只在方法中才生效
func (man Man) introduce(){

	fmt.Println("Hello My name is ",man.name, ", i am " , man.age , " years old" );
}


// 指针接收者的方法可以修改接收者指向的值（就像 Scale 在这做的）。由于方法经常需要修改它的接收者，指针接收者比值接收者更常用。
// 对数据的修改会全局生效
func (man *Man) introduce2(){

	fmt.Println("Hello My name is ",man.name, ", i am " , man.age , " years old" );
	man.name="jay2";
	man.age=12;
}


func main() {
	
	man := Man{"Jay",10};
	
	man.introduce();
	
   fmt.Println("Hello, World!")
}
