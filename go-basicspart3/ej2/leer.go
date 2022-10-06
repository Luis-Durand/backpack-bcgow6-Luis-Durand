package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	data, err := os.ReadFile("./test.csv")

	if err != nil {
		fmt.Println(err)

	}

	fmt.Println(strings.ReplaceAll(string(data), ";", "\t\t\t"))

}
