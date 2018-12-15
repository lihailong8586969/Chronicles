type Rectangle struct{

	leftX float64;
	topY float64;
	heigh float64;
	width float64;
}


func (rect *Rectangle) area() float64{

	return rect.width * rect.heigh;
}
