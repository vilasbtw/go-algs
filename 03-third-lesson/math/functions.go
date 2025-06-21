package math

// Named Return Values
func DoubleNumber(number int) (total int) {
	total = number * 2

	// not required to write the returned variable name
	return
}

// Variadic Function (of integers)
func SumNumbers(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}
	return total
}

/*
func Quadradic(numbers ...int) int {
	return SumNumbers(numbers...) * SumNumbers(numbers...)
}
*/

// High-order function
func QuadradicOfTheTotal(numbers ...int) func() int {

	result := func() int {
		return SumNumbers(numbers...) * SumNumbers(numbers...)
	}

	return result
}
