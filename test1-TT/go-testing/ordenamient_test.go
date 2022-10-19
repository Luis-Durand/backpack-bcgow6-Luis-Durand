package testing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenamiento(t *testing.T) {

	sliceNumbers := []int{5, 4, 3, 2, 1}
	expectedSlice := []int{1, 2, 3, 4, 5}
	result := Ordenamiento(sliceNumbers)

	assert.Equal(t, expectedSlice, result, "se esperaba %v y dio %v", expectedSlice, result)

}
