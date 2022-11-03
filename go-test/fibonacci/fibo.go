package fibonacci

func Fibonacci(num int) int {
	a := 0
	b := 1
	c := num

	for i := 2; i <= num; i++ {
		c = a + b
		a = b
		b = c
	}
	return c

}
