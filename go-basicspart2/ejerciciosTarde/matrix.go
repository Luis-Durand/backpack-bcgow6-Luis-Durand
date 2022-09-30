package main

import "fmt"

type Matrix struct {
	Alto        float64
	Ancho       float64
	Cuadratica  bool
	ValorMaximo float64
}

func (m *Matrix) Set(alto, ancho, max float64, cuadratica bool) {
	m.Alto = alto
	m.Ancho = ancho
	m.Cuadratica = cuadratica
	m.ValorMaximo = max

}

func (m *Matrix) Print() {

	fmt.Printf("Alto: %fm\nAncho: %fm\nCuadratica: %t\nValorMaximo: %fm\n", m.Alto, m.Ancho, m.Cuadratica, m.ValorMaximo)
}

func Matrixx() {

	newMatrix := Matrix{}
	newMatrix.Set(3002, 100, 5000, true)
	newMatrix.Print()

}
