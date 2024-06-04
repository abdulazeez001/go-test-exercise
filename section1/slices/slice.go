package slices

func isValEven(value int) bool {
	return value%2 == 0
}

func SumOfAllEvenNum(listOfNums []int) int {

	sum := 0

	if listOfNums == nil {
		return sum
	}

	for _, value := range listOfNums {
		if isValEven(value) {
			sum += value
		}
	}

	return sum
}
