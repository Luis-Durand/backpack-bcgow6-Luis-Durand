/*
	Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento.

Para ello requieren realizar un programa que se encargue de calcular el precio total de Productos, Servicios y Mantenimientos. Debido a la fuerte demanda y para optimizar la velocidad requieren que el cálculo de la sumatoria se realice en paralelo mediante 3 go routines.

Se requieren 3 estructuras:
Productos: nombre, precio, cantidad.
Servicios: nombre, precio, minutos trabajados.
Mantenimiento: nombre, precio.

Se requieren 3 funciones:
Sumar Productos: recibe un array de producto y devuelve el precio total (precio * cantidad).
Sumar Servicios: recibe un array de servicio y devuelve el precio total (precio * media hora trabajada, si no llega a trabajar 30 minutos se le cobra como si hubiese trabajado media hora).
Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.

Los 3 se deben ejecutar concurrentemente y al final se debe mostrar por pantalla el monto final (sumando el total de los 3).
*/
package main

import "fmt"

type Productos struct {
	Nombre   string
	Precio   int
	Cantidad int
}

type Servicios struct {
	Nombre  string
	Precio  int
	Minutos int
}

type Mantenimiento struct {
	Nombre string
	Precio int
}

func sumProd(products []Productos, c chan int) {
	totPrecio := 0
	totCantidad := 0

	for _, i := range products {
		totPrecio += i.Precio
		totCantidad += i.Cantidad

	}
	c <- totPrecio * totCantidad
}

func sumServ(servicio []Servicios, c chan int) {
	totPrecio := 0
	totMinutos := 0

	for _, i := range servicio {
		totMinutos += i.Minutos
		totPrecio += i.Precio
	}

	if totMinutos%30 < 0 {
		totMinutos = (totMinutos / 30) + 1
		c <- totPrecio * totMinutos

	}

	c <- totPrecio * (totMinutos / 30)
}

func sumMant(mant []Mantenimiento, c chan int) {
	total := 0
	for _, i := range mant {
		total += i.Precio
	}
	c <- total
}

func main() {
	Total := 0
	canal := make(chan int)

	helado := Productos{"Helado", 70, 200}
	vasos := Productos{"Vasos", 10, 500}
	arrProd := []Productos{helado, vasos}

	limpieza := Servicios{"Limpieza", 300, 100}
	cocinar := Servicios{"cocinar", 100, 100}
	arrServ := []Servicios{limpieza, cocinar}

	polvo := Mantenimiento{"Polvo", 200}
	compu := Mantenimiento{"Compu", 100}
	arrMant := []Mantenimiento{polvo, compu}

	go sumProd(arrProd, canal)
	totProd := <-canal

	go sumServ(arrServ, canal)
	totServ := <-canal

	go sumMant(arrMant, canal)
	totMant := <-canal

	Total += totProd + totServ + totMant
	fmt.Println(Total)
}
