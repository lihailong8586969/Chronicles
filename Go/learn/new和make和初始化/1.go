// new 和 make 的区别


// Go语言中的内建函数new和make是两个用于内存分配的原语（allocation primitives）。
// 对于初学者，这两者的区别也挺容易让人迷糊的。
// 简单的说，new只分配内存，make用于slice，map，和channel的初始化。

// 1. new
// 这是一个用来分配内存的内建函数，但是与C++不一样的是，它并不初始化内存，只是将其置零。
// 也就是说，new(T)会为T类型的新项目，分配被置零的存储，并且返回它的地址，
// 一个类型为*T的值。在Go的术语中，其返回一个指向新分配的类型为T的指针，
// 这个指针指向的内容的值为零（zero value）。注意并不是指针为零。
// Go语言中的对象没有C++中的构造函数，如果用C来描述，Go中的new大概相当于：
T *t = (T*)malloc(sizeof(T))
memset(t, 0, sizeof(T))

// 其实，上面的描可能也不是很准确，也许用*t=zerovalue更准确。
// 因为对于不同的数据类型，零值的意义是完全不一样的。
// 比如，对于bool类型，零值为false；int的零值为0；string的零值是空字符串：
b := new(bool)

fmt.Println(*b)

i := new(int)

fmt.Println(*i)

s := new(string)

fmt.Println(*s)
