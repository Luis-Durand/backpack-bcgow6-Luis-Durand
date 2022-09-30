package main

import "fmt"

type Alumnos struct {
	Nombre   string
	Apellido string
	Dni      int
	Fecha    string
}

func (a Alumnos) detalle(nombre, apellido, fecha string, dni int) {

	fmt.Printf("Nombre: %s\nApellido: %s\nDni: %d\nFecha: %s\n", nombre, apellido, dni, fecha)

}

func Meme() {

	Juan := Alumnos{}
	Juan.detalle("Juan", "Perez", "10/05", 23123)

}
