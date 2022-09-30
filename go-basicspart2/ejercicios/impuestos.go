package ejercicios

func Impuestos(sueldo int) int {

	if sueldo > 150000 {

		return sueldo - (sueldo*27)/100
	} else {
		return sueldo - (sueldo*10)/100
	}

}
