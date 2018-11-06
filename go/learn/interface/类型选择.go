package main

import "fmt"

func do(i interface{}) {

	// switch v := i.(type) {
	// 	case T:
	// 	    // v 的类型为 T
	// 	case S:
	// 	    // v 的类型为 S
	// 	default:
	// 	    // 没有匹配，v 与 i 的类型相同
	// }

	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
