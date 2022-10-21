package testing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividir(t *testing.T) {

	num := 10
	den := 2
	expectedResult := 5
	result, err := Dividir(num, den)

	assert.Nil(t, err, "El denominador no puede ser 0")

	assert.Equal(t, expectedResult, result, "Se esperaba %d y nos dio %d", expectedResult, result)

}
