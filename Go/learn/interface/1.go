package main

import (
	"fmt"
	"reflect"
)

type Phone interface {

	call()

	charge()
}



type NokiaPhone struct {
}
func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}
func (nokiaPhone NokiaPhone) charge() {
	fmt.Println("I am Nokia, I'm charging")
}



type IPhone struct {
}
func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}
func (iPhone IPhone) charge(){
	fmt.Println("I am iPhone, I'm charging")
}




type Man struct{

	name string
}


// Go程序抽象的基本原则依赖于接口而不是实现，优先使用组合而不是继承。
func main() {

	var phone Phone



	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

	fmt.Println( "type : ",  reflect.TypeOf(phone) )


}