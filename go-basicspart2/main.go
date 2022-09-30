package main

import (
	"fmt"

	"github.com/Luis-Durand/backpack-bcgow6-Luis-Durand/go-basicspart2/ejercicios"
)

func main() {

	fmt.Println(ejercicios.Impuestos(200000))

	promedio, error := ejercicios.Promedio(10, 5, 3, 4)

	if promedio != 0 {
		fmt.Println("El promedio es ", promedio)

	} else {
		fmt.Println(error)
	}

	ejercicios.Estadistica()
	ejercicios.Cantidad()
}
