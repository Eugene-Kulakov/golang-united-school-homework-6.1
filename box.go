package golang_united_school_homework

import (
	"errors"
)

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
	if len(b.shapes) >= b.shapesCapacity {
		return errors.New("unable to add shape, box is full")
	}

	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i >= len(b.shapes) {
		return nil, errors.New("index out of the range")
	}
	//if b.shapes[i] == nil {
	//	return nil, errors.New("shape by index doesn't exist")
	//}

	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if i >= len(b.shapes) {
		return nil, errors.New("index out of the range")
	}
	if b.shapes[i] == nil {
		return nil, errors.New("shape by index doesn't exist")
	}

	ans := b.shapes[i]
	copy(b.shapes[i:], b.shapes[i+1:])
	b.shapes[len(b.shapes)-1] = nil
	b.shapes = b.shapes[:len(b.shapes)-1]
	return ans, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if i >= len(b.shapes) {
		return nil, errors.New("index out of the range")
	}
	if b.shapes[i] == nil {
		return nil, errors.New("shape by index doesn't exist")
	}

	ans := b.shapes[i]
	b.shapes[i] = shape
	return ans, nil

}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var ans float64
	for _, shape := range b.shapes {
		if shape != nil {
			ans += shape.CalcPerimeter()
		}
	}
	return ans
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var ans float64
	for _, shape := range b.shapes {
		if shape != nil {
			ans += shape.CalcArea()
		}
	}
	return ans
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	success := false
	var newShapes []Shape
	for _, shape := range b.shapes {
		_, ok := shape.(*Circle)
		if !ok {
			newShapes = append(newShapes, shape)
		} else {
			success = true
		}
	}
	if !success {
		return errors.New("circles are not exist in the list")
	}
	b.shapes = newShapes
	return nil
}
