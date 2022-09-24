package ejercicios

import "fmt"

func Edad() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	employees["Federico"] = 25

	var edades int

	for key, element := range employees {

		if key == "Benjamin" {
			fmt.Printf("Edad de Benjamin: %v \n", element)
		}
		if element > 21 {
			edades++
		}

	}
	fmt.Printf("Mis empleados mayores a 21 son : %v \n", edades)
	fmt.Println(employees)
	delete(employees, "Pedro")

}
