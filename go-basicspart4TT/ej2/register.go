package main

import (
	"fmt"
	"os"
)

type Clientes struct {
	Legajo    int
	Nombre    string
	Apellido  string
	Dni       int
	Teléfono  int
	Domicilio string
}

var id int

func createId() int {
	id++
	return id
}

func readArch() {

	defer func() {
		err := recover()

		if err != nil {
			fmt.Println("error: el archivo indicado no fue encontrado o está dañado")
		}

	}()

	_, err := os.ReadFile("Customers.txt")

	if err != nil {
		panic(err)
	}

}

func main() {

	data := createId()

	if data == 0 {
		panic("Error en la funcion")
	}
	readArch()
	fmt.Println("La ejecucion siguio")
}
