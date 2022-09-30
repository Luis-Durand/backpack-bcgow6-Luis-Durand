package ejercicios

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func minimumOp(value ...int) float64 {

	result := value[0]
	for _, v := range value {
		if v < 0 {
			return 0
		}
		if v < result {
			result = v
		}
	}
	return float64(result)
}

func averageOp(value ...int) float64 {

	total := 0

	for _, v := range value {
		if v < 0 {
			return 0
		}
		total += v

	}

	return float64(total) / float64(len(value))

}

func maximumOp(value ...int) float64 {

	result := value[0]

	for _, v := range value {
		if v < 0 {
			return 0
		}
		if v > result {
			result = v
		}
	}
	return float64(result)

}

func operation(calculo string) (func(...int) float64, error) {

	switch calculo {
	case minimum:
		return minimumOp, nil
	case average:
		return averageOp, nil
	case maximum:
		return maximumOp, nil
	default:
		return nil, errors.New("Invalid operation")
	}

}

func Estadistica() {

	minFunc, err := operation(minimum)
	if err != nil {
		fmt.Println(err)
		return
	}
	averageFunc, err := operation(average)
	if err != nil {
		fmt.Println(err)
		return
	}
	maximum, err := operation(maximum)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("El minimo es %.2f\n", minFunc(10, 20, 23, 123))
	fmt.Printf("El promedio es %.2f\n", averageFunc(5, 20, 23, 123))
	fmt.Printf("El maximo es %.2f\n", maximum(300, 20, 23, 123))

}
