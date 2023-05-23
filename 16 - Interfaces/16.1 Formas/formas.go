package main

import "fmt"

type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A área da forma é %0.2f\n", f.area())
}

type rectangle struct {
	altura  float64
	largura float64
}

func (r rectangle) area() float64 {
	return r.altura * r.largura
}

type circle struct {
	raio float64
}

func (c circle) area() float64 {
	return c.raio * c.raio * 3.14
}

func main() {
	fmt.Println("Interfaces")

	r := rectangle{10, 15}
	escreverArea(r)
}
