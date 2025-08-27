package main

import "fmt"

type XFloat32 float32

func (f XFloat32) XPrint() {
	fmt.Println(f)
}

// func (f float32) YPrint() { 		// error because float32 is not from the same package
// 	fmt.Println(f)
// }

type Vertex struct {
	X, Y float32
}

func (v *Vertex) ScaleMethod(f float32) {
	v.X *= f
	v.Y *= f
}

func ScaleFunction(v *Vertex, f float32) {
	v.X *= f
	v.Y *= f
}

func methods_func() {
	num := XFloat32(2.1)
	num.XPrint()

	v := Vertex{3, 4}
	v.ScaleMethod(2)      //interpreted automatically as pointer, i.e. (&v).ScaleMethod(2)
	ScaleFunction(&v, 10) //must be explicit as pointer
	fmt.Println(v)

	p := &v //pointer to a vertex
	p.ScaleMethod(0.5)
	ScaleFunction(p, 0.1)
	fmt.Println(v)
}
