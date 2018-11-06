


// switch 是编写一连串 if - else 语句的简便方法。它运行第一个值等于条件表达式的 case 语句。
// Go 的 switch 语句类似于 C、C++、Java、JavaScript 和 PHP 中的，
// 不过 Go 只运行选定的 case，而非之后所有的 case。 
// 实际上，Go 自动提供了在这些语言中每个 case 后面所需的 break 语句。 
// 除非以 fallthrough 语句结束，否则分支会自动终止。 
// Go 的另一点重要的不同在于 switch 的 case 无需为常量，且取值不必为整数。

package main

import (
	"fmt"
	"runtime"
)

func main() {
	
	fmt.Print("Go runs on ")

	// 同 for 一样， if 语句可以在条件表达式前执行一个简单的语句。 
	switch os := runtime.GOOS; os {
		
		// 不用加 break
		// switch 的 case 语句从上到下顺次执行，直到匹配成功时停止。
		//（例如，
		// switch i {
		//	case 0:
		//	case f():
		// }
		// 在 i==0 时 f 不会被调用。）
		case "darwin":
		
			fmt.Println("OS X.")
		
		case "linux":
		
			fmt.Println("Linux.")
		
		default:
			// freebsd, openbsd,
			// plan9, windows...
			fmt.Printf("%s.", os)
	}


	// 没有条件的 switch, 没有条件的 switch 同 switch true 一样。 
	// 这种形式能将一长串 if-then-else 写得更加清晰。
	// 相当于 if (else if) (else if) (else if) else
	i:=10;
	
	switch {
	case i < 12:
		fmt.Println("比12小")
	case i < 17:
		fmt.Println("比17小")
	default:
		fmt.Println("Good evening.")
	}
	// 输出 : 比12小

}
