package testing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {

	firstNum := 10
	secondNum := 5
	expectedResult := 5
	result := Restar(firstNum, secondNum)

	assert.Equal(t, expectedResult, result, "Se esperaba %d y salio %d", expectedResult, result)

}
