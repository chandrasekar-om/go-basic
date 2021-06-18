package aggregateType

import "fmt"

type Vertex struct {
	X int
	Y int
}

func AccessFieldsWithDot() {
	v := Vertex{1, 3}
	v.X = 4
	fmt.Println(v)
}

func AccessFieldsWithPointer() {
	v := Vertex{2, 4}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

func AccessFieldsWithLiterals() {
	v := Vertex{1, 2}
	v1 := Vertex{X: 3}
	v2 := Vertex{}
	p := &Vertex{4, 6}
	fmt.Println(v, p, v1, v2)
}

func newVertex(x, y int) *Vertex {
	p := Vertex{x, y}
	return &p
}

func EncapsulateNewStructCreation() {
	v := newVertex(100, 200)
	fmt.Println(v)
}
