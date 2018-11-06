package main

// 打包导入
import (

	"fmt"
	"math"
)

// fmt 是 format 的缩写

// 依次导入
// import "fmt"
// import "math"


// 包级别结构体
type Vertex struct {
	X int
	Y int
}


// 包级别变量
var c, python, java bool
var i, j int = 1, 2

// 这个例子演示了具有不同类型的变量。 同时与导入语句一样，变量的定义“打包”在一个语法块中。
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// 包级别全局常量
const (
	Big   = 1 << 100
	Small = Big >> 99
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

	// 函数级别结构体
	type Vertex struct {
		X int
		Y int
	}
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	// 结构体文法
    // 结构体文法表示通过结构体字段的值作为列表来新分配一个结构体。
	// 使用 Name: 语法可以仅列出部分字段。（字段名的顺序无关。）
	// 特殊的前缀 & 返回一个指向结构体的指针。
	v1 = Vertex{1, 2}  // 类型为 Vertex
	v2 = Vertex{X: 1}  // Y:0 被省略
	v3 = Vertex{}      // X:0 和 Y:0
	p  = &Vertex{1, 2} // 类型为 *Vertex

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
	i, j := 42, 2701
	

	var a = "string";
	var b string = "runoob.com";
	var c bool ;
	fmt.Println("Hello, playground");

	// 输出 : v is of type int and value = 98 
	// := 是定义并赋值
	// =  是重新赋值
	v := 42 // 定义变量 v 并赋值
	v = 98 // 重新给 v 赋值
	fmt.Printf("v is of type %T and value = %d \n", v,v)

	// 常量的定义与变量类似，只不过使用 const 关键字。
	// 常量可以是字符、字符串、布尔或数字类型的值。
	// 常量不能使用 := 语法定义。
	const World = "世界"

	// 函数级别局部常量
	const (
		Big   = 1 << 100
		Small = Big >> 99
	)


	// if 语句除了没有了 `( )` 之外（甚至强制不能使用它们），看起来跟 C 或者 Java 中的一样，而 `{ }` 是必须的。（耳熟吗？）
	if 1<9 {

  		fmt.Println("the spice must flow");
	}

	// if 的便捷语句,跟 for 一样，`if` 语句可以在条件之前执行一个简单的语句, 由这个语句定义的变量的作用域仅在 if 范围之内。
	if v := math.Pow(x, n); v < lim {
		
		return v
	}

	// if 和 else, 在 if 的便捷语句定义的变量同样可以在任何对应的 else 块中使用。
	if v := math.Pow(x, n); v < lim {
	
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}

	// defer , defer 语句会延迟函数的执行直到上层函数返回。 延迟调用的参数会立刻生成，但是在上层函数返回前函数都不会被调用。
	defer fmt.Println("world")

	// defer 栈 , 延迟的函数调用被压入一个栈中。当函数返回时， 会按照后进先出的顺序调用被延迟的函数调用。 阅读博文了解更多关于 defer 语句的信息。
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}


	// switch 一个结构体（`struct`）就是一个字段的集合。 除非以 fallthrough 语句结束，否则分支会自动终止。
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}


	// 没有条件的 switch, 没有条件的 switch 同 `switch true` 一样。 这一构造使得可以用更清晰的形式来编写长的 if-then-else 链。
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}


	// 在导入了一个包之后，就可以用其导出的名称来调用它。在 Go 中，首字母大写的名称是被导出的。Foo 和 FOO 都是被导出的名称。名称 foo 是不会被导出的。
	fmt.Println(math.Pi);


	fmt.Println(add(42, 13));


	a, b := swap("hello", "world");
	fmt.Println(a, b);

	fmt.Println(split(17));


	// Go 只有一种循环结构——`for` 循环。
	// 基本的 for 循环除了没有了 `( )` 之外（甚至强制不能使用它们），看起来跟 C 或者 Java 中做的一样，而 `{ }` 是必须的。
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)


	// 跟 C 或者 Java 中一样，可以让前置、后置语句为空。
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)

	// for 是 Go 的 “while”, 基于此可以省略分号：C 的 while 在 Go 中叫做 `for`。
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// 死循环 : 如果省略了循环条件，循环就不会结束，因此可以用更简洁地形式表达死循环。
	for{

	}




	// 指针
	i, j := 42, 2701
	
	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j




}

