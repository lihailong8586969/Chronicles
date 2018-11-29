func main() {

	var phone Phone



	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

	// 打印变量类型
	fmt.Println( "type : ",  reflect.TypeOf(phone) )


}