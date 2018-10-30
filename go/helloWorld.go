package main

// 打包导入
import (

	"fmt"
	"math"
)

// 依次导入
// import "fmt"
// import "math"

// 包级别变量
var c, python, java bool
var i, j int = 1, 2

// 这个例子演示了具有不同类型的变量。 同时与导入语句一样，变量的定义“打包”在一个语法块中。
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// 函数
func add(x int, y int) int {

	return x + y
}

// 当两个或多个连续的函数命名参数是同一类型，则除了最后一个类型之外，其他都可以省略。
func add(x, y int) int {
	return x + y
}

// 函数多值返回
func swap(x, y string) (string, string) {

	return y, x
}

// 没有参数的 return 语句返回结果的当前值。也就是`直接`返回。
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {


	// 函数级别变量
	var c, python, java bool
	var c, python, java = true, false, "no!"


	// 变量在定义时没有明确的初始化时会赋值为_零值_。
	// 零值是： 数值类型为 `0`， 布尔类型为 `false`， 字符串为 `""`（空字符串）。
	var i int
	var f float64
	var b bool
	var s string

	// 在函数中，`:=` 简洁赋值语句在明确类型的地方，可以用于替代 var 定义。
	// 函数外的每个语句都必须以关键字开始（`var`、`func`、等等），`:=` 结构不能使用在函数外。
	k := 3
	c, python, java := true, false, "no!"


	var a = "string";
	var b string = "runoob.com";
	var c bool ;
	fmt.Println("Hello, playground");

	// 输出 : v is of type int and value = 98 
	v := 42 // 定义变量 v 并赋值
	v = 98 // 重新给 v 赋值
	fmt.Printf("v is of type %T and value = %d \n", v,v)

	// 常量的定义与变量类似，只不过使用 const 关键字。

	// 常量可以是字符、字符串、布尔或数字类型的值。
	// 常量不能使用 := 语法定义。
	const World = "世界"

	// 可以直接省略括号
	if (1<9) {

  		fmt.Println("the spice must flow");
	}

	// 在导入了一个包之后，就可以用其导出的名称来调用它。在 Go 中，首字母大写的名称是被导出的。Foo 和 FOO 都是被导出的名称。名称 foo 是不会被导出的。
	fmt.Println(math.Pi);


	fmt.Println(add(42, 13));


	a, b := swap("hello", "world");
	fmt.Println(a, b);

	fmt.Println(split(17));


}

