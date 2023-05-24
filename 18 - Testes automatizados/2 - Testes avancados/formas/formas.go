package formas

import "math"

// Forma é uma interface que descreve uma forma geométrica
type Forma interface {
	area() float64
}

// Rectangle representa um retângulo
type Rectangle struct {
	Altura  float64
	Largura float64
}

// Area calcula a área do retângulo
func (r Rectangle) Area() float64 {
	return r.Altura * r.Largura
}

// Circle representa um círculo
type Circle struct {
	raio float64
}

// Area calcula a área do círculo
func (c Circle) Area() float64 {
	return c.raio * c.raio * math.Pi
}
