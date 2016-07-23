package interfaces

type Shape interface {
	Area() int
}

type Square struct{
	side int
}

func(aSquare Square) Area() int{
	return aSquare.side * aSquare.side
}

type Rectangle struct{
	length int
	breadth int
}

func(aRectangle Rectangle) Area() int{
	return aRectangle.length * aRectangle.breadth
}

type Hybrid struct{
	square Square
	rectangle Rectangle
}

func(aHybrid Hybrid) Area() int{
	return aHybrid.square.Area() + aHybrid.rectangle.Area()
}