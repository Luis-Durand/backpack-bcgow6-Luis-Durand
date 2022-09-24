package ejercicios

import (
	"fmt"
	"strings"
)

func Letras() {

	var palabra string = "auto"
	large := len(palabra)
	fmt.Printf("El largo de la palabra es de %v \n", large)

	newData := strings.Split(palabra, "")

	for i := 0; i < len(newData); i++ {
		fmt.Print(newData[i] + "-")

		if i == 3 {
			fmt.Print("\n")
		}
	}
}
