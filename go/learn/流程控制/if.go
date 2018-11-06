package main

import (
	"fmt"
)



func pow(x, n, lim float64) float64 {

	// 同 for 一样， if 语句可以在条件表达式前执行一个简单的语句。 
	// 该语句声明的变量作用域仅在 if 之内。（可以在最后的 return 语句处使用 v 看看。）
	if v := math.Pow(x, n); v < lim {

		return v
	}else{

		// 在 if 的简短语句中声明的变量同样可以在任何对应的 else 块中使用。


	}
	return lim
}


func main() {
	
	i:=10;
	if i>5 {

		fmt.Println("大");
	}

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

}
