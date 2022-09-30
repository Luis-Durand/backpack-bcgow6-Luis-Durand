package main

import (
	"fmt"
	"os"
)

type Producto struct {
	Id       int
	Precio   int
	Cantidad int
}

func Archivo() {

	firstProduct := Producto{
		Id:       1,
		Precio:   200,
		Cantidad: 1000,
	}
	secondProduct := Producto{
		Id:       2,
		Precio:   500,
		Cantidad: 400,
	}
	thirdProduct := Producto{
		Id:       3,
		Precio:   700,
		Cantidad: 1200,
	}

	msg := fmt.Sprintf("Id;Precio;Cantidad\n%d;%d;%d\n%d;%d;%d\n%d;%d;%d\n", firstProduct.Id, firstProduct.Precio, firstProduct.Cantidad, secondProduct.Id, secondProduct.Precio, secondProduct.Cantidad, thirdProduct.Id, thirdProduct.Precio, thirdProduct.Cantidad)

	os.WriteFile("./test.csv", []byte(msg), 0644)

}

func main() {

	Archivo()
}
