package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Go does not have classes. However, you can define methods on types.
// Remember: a method is just a function with a receiver argument.

// A method is a function with a special receiver argument.
// (v Vertex) 是 Abs 方法的接收器
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}


// Here's Abs written as a regular function with no change in functionality.
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}



func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
