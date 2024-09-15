package main

import "fmt"

type Shape interface {
	Draw()
}

type Circle struct{}

func (c Circle) Draw() {
	fmt.Println("Drawing Circle")
}

type Square struct{}

func (s Square) Draw() {
	fmt.Println("Drawing Square")
}

func ShapeFactory(shapeType string) Shape {
	switch shapeType {
	case "circle":
		return Circle{}
	case "square":
		return Square{}
	default:
		return nil
	}
}

func main() {
	shape1 := ShapeFactory("circle")
	shape1.Draw()

	shape2 := ShapeFactory("square")
	shape2.Draw()
}
