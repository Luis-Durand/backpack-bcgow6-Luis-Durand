/*
	Una importante empresa de ventas web necesita agregar una funcionalidad para agregar productos a los usuarios. Para ello requieren que tanto los usuarios como los productos tengan la misma dirección de memoria en el main del programa como en las funciones.

Se necesitan las estructuras:
Usuario: Nombre, Apellido, Correo, Productos (array de productos).
Producto: Nombre, precio, cantidad.
Se requieren las funciones:
Nuevo producto: recibe nombre y precio, y retorna un producto.
Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el producto al usuario.
Borrar productos: recibe un usuario, borra los productos del usuario.
*/
package main

import "fmt"

type Usuario struct {
	Nombre    string
	Apellido  string
	Correo    string
	Productos []producto
}

type producto struct {
	Nombre   string
	Precio   float64
	Cantidad float64
}

func nuevoProducto(nombre string, precio float64) producto {
	return producto{nombre, precio, 0}

}

func agregarProducto(usuario *Usuario, p *producto, cantidad float64) {
	newProduct := producto{p.Nombre, p.Precio, cantidad}

	usuario.Productos = []producto{newProduct}

}

func borrarProducto(usuario *Usuario) {
	usuario.Productos = []producto{}

}

func main() {

	datita := producto{"alberto", 23.4324, 23.3422}

	data := []producto{datita}

	juan := Usuario{"Juan", "Meno", "esae@gmai.com", data}
	/* 	borrarProducto(&juan)
	 */
	/*  agregarProducto(&juan, &datita, 200) */

	newProd := nuevoProducto("Alfajor", 500)
	fmt.Println(juan)
	fmt.Println(newProd)

}
