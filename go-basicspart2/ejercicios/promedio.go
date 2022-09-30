package ejercicios

import "errors"

func Promedio(calificaciones ...int) (prom float64, err error) {
	var resultado float64

	for _, value := range calificaciones {

		if value < 1 {
			err = errors.New("Calificacion no puede ser menor a 1")
			return
		}

		resultado += float64(value)

	}
	prom = resultado / float64(len(calificaciones))
	return

}
