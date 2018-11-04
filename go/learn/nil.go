if err != nil {
    // do something....
}

// 当出现不等于nil的时候，说明出现某些错误了，
// 需要我们对这个错误进行一些处理，
// 而如果等于nil说明运行正常。
// 那什么是nil呢？查一下词典可以知道，
// nil的意思是无，或者是零值。
// 零值，zero value，是不是有点熟悉？
// 在Go语言中，如果你声明了一个变量但是没有对它进行赋值操作，
// 那么这个变量就会有一个类型的默认零值。这是每种类型对应的零值：

// bool      -> false                              
// numbers -> 0                                 
// string    -> ""      

// pointers -> nil
// slices -> nil
// maps -> nil
// channels -> nil
// functions -> nil
// interfaces -> nil



// Go的文档中说到，nil是预定义的标识符，代表指针、通道、函数、接口、映射或切片的零值，也就是预定义好的一个变量：