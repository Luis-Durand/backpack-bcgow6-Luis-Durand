package main

import (
	"fmt"
	"os"
)

func clientes() {

	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}
	}()

	_, err := os.ReadFile("customers.txt")
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}

}

func main() {

	clientes()
	fmt.Println("ejecucion finalizada")

}
