// import 原理 : https://www.imooc.com/video/16850

// 如果一个 main 包导入其他包，包被顺序导入

// 如果导入的包中有依赖其他包( 包 B ) , 
// 会首先导入 B 包 , 然后初始化 B 包中常量和变量 ,
// 最后如果 B 包中有init, 会自动执行 init();

// 所有包导入完后, 才敢于对于 main 包中常量和变量进行初始化 , 
// 然后执行 main 包中的 init 函数( 如果存在 ), 
// 最后才会执行 main 函数

// 每个 Go 语言源代码文件开头都必须有一个 package 声明, 表示源码文件所属文件包

// 必须有 main 包 , 包名必须是 main ,main 包里面还必须有 main 函数

// 同一个路径下只能存在一个 package , 一个 package 可以拆成多个源文件组成

// package 包名要和目录名一样( 规范 )

// 不得导入源码中未使用到的 package

// 如果一个包被导入多次，则这个包只会被导入 1 次

import{
	
	// 还能直接用 url 写
	socketio "https://github.com/googollee/go-socket.io"
}