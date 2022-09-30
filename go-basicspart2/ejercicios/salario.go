package ejercicios

func Salario(minTrabajados int, categoria string) (total int, err error) {

	switch categoria {

	case "A":
		total = (minTrabajados / 60) * 3000
		return
	case "B":
		total = (minTrabajados / 60) * 1500
	case "C":
		total = (minTrabajados / 60) * 1000
	}

	return
}
