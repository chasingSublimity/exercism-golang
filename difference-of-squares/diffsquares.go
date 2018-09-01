package diffsquares

// Difference returns the difference between the outputs of
// SquareOfSums and SumOfSquares
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}

// SquareOfSums returns the square of the
// sum of all natural numbers between 0 and n
func SquareOfSums(n int) int {
	var sum int
	for i := 0; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

// SumOfSquares returns the sum of
// all squared natural numbers between 0 and n
func SumOfSquares(n int) int {
	var sum int
	for i := 0; i <= n; i++ {
		sum += i * i
	}
	return sum
}
