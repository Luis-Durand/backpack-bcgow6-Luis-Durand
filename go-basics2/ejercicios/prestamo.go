package ejercicios

import "fmt"

func Prestamo() {

	var edad int = 40
	var empleado bool = true
	var año int = 1
	var sueldo int = 5000

	if edad > 22 {
		fmt.Println("Es mayor a 22")
		if empleado {
			fmt.Println("Tiene trabajo")

			if año > 1 {
				fmt.Println("Tiene mas d un año")

				if sueldo > 100000 {

					fmt.Println("No se te cobra impuesto")
				} else {
					fmt.Println("Se le cobra impuesto")
				}

			} else {
				fmt.Println("No tiene mas d un año d laburo")
			}

		} else {
			fmt.Println("No tiene laburo")
		}
	} else {
		fmt.Println("No sos mayor a 22")
	}

}
