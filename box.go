package golang_united_school_homework

import "fmt"

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	shapesLength := len(b.shapes)
	if shapesLength == b.shapesCapacity {
		return fmt.Errorf("Out of the shapesCapacity range")
	} else {
		b.shapes = append(b.shapes, shape)
		return nil
	}
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	err := isOutOfRange(i, len(b.shapes))
	if err != nil {
		return nil, err
	}

	var shape Shape
	shape = b.shapes[i]
	return shape, nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	err := isOutOfRange(i, len(b.shapes))
	if err != nil {
		return nil, err
	}

	var shape Shape
	shape = b.shapes[i]
	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
	return shape, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {

	err := isOutOfRange(i, len(b.shapes))
	if err != nil {
		return nil, err
	}
	var replacedShape Shape
	replacedShape = b.shapes[i]
	b.shapes[i] = shape
	return replacedShape, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	sumPerimeter := 0.0
	for _, shape := range b.shapes {
		sumPerimeter += shape.CalcPerimeter()
	}
	return sumPerimeter
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	sumArea := 0.0
	for _, shape := range b.shapes {
		sumArea += shape.CalcArea()
	}
	return sumArea
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	shapes := make([]Shape, 0, len(b.shapes))
	for _, shape := range b.shapes {
		_, ok := shape.(*Circle)
		if !ok {
			shapes = append(shapes, shape)
		}
	}
	if len(shapes) == len(b.shapes) {
		return fmt.Errorf("Circles are not exist in the list")
	}
	b.shapes = shapes
	return nil
}

func isOutOfRange(index, length int) error {
	if index < 0 || index >= length {
		err := fmt.Errorf("Index %d is out of the range", index)
		return err
	}
	return nil
}
