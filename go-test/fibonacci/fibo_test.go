package fibonacci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {

	numberFibo := Fibonacci(0)
	numberFibo1 := Fibonacci(1)
	numberFibo2 := Fibonacci(2)
	numberFibo3 := Fibonacci(3)
	numberFibo10 := Fibonacci(10)

	assert.Equal(t, 0, numberFibo)
	assert.Equal(t, 1, numberFibo1)
	assert.Equal(t, 1, numberFibo2)
	assert.Equal(t, 2, numberFibo3)
	assert.Equal(t, 55, numberFibo10)

}
