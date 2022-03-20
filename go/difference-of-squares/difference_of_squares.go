package diffsquares

func SquareOfSum(n int) int {
	//panic("Please implement the SquareOfSum function")
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

func SumOfSquares(n int) int {
	panic("Please implement the SumOfSquares function")
}

func Difference(n int) int {
	panic("Please implement the Difference function")
}
