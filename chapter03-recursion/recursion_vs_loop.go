package chapter03_recursion

func SumLoop(x int) int {
	sum := 0
	for ; x > 0; x-- {
		sum += x
	}

	return sum
}

func SumRecursion(x int) int {
	if x == 0 {
		return 0
	}

	return x + SumRecursion(x-1)
}

func GaussFormula(x int) int {
	return x * (x + 1) / 2
}
