package main

import "fmt"

type Man struct{

	name string;
	age int;
}

// 对数据的修改只在方法中才生效
func (man Man) introduce(){

	fmt.Println("Hello My name is ",man.name, ", i am " , man.age , " years old" );
}

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
