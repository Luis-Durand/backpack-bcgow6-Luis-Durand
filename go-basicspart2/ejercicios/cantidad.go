/* Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas. Por el momento solo tienen tarántulas, hamsters, perros, y gatos, pero se espera que puedan haber muchos más animales que refugiar.

perro necesitan 10 kg de alimento
gato 5 kg
Hamster 250 gramos.
Tarántula 150 gramos.

Se solicita:
Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado y que retorne una función y un mensaje (en caso que no exista el animal)
Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.
*/

package ejercicios

import (
	"errors"
	"fmt"
)

func Perro(perros int) int {
	return perros * 10

}
func Gato(gatos int) int {
	return gatos * 5

}
func Hamster(hamsters int) int {
	return hamsters * 250

}
func Tarantula(tarantulas int) int {
	return tarantulas * 150

}

func Animal(valor string) (func(int) int, error) {

	switch valor {
	case "perro":
		return Perro, nil
	case "gato":
		return Gato, nil
	case "tarantula":
		return Tarantula, nil
	case "hamster":
		return Hamster, nil
	default:
		return nil, errors.New("No existe ese animal")
	}

}

func Cantidad() {
	dog, err := Animal("perro")

	if err != nil {
		fmt.Println(err)
		return
	}

	cat, err := Animal("gato")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Para el perro hay q comprar %d kg\n", dog(5))
	fmt.Printf("Para el gato hay q comprar %d kg\n", cat(5))

}
