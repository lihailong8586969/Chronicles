// 变量关联的存储在函数返回之后依然存在



type Rect struct {

	x, y float64

	width, height float64

}



rect3 := &Rect{0, 0, 100, 200}

rect4 := &Rect{width: 100, height: 200}



func NewRect(x, y, width, height float64) *Rect {

	// 局部变量关联的内存不会被释放
	return &Rect{x, y, width, height}
}

