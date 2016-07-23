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
	aShape Shape
	anotherShape Shape
}

func(aHybrid Hybrid) Area() int{
	return aHybrid.aShape.Area() + aHybrid.anotherShape.Area()
}