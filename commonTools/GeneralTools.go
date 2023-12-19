package commonTools

func PrintErr(err error) {
	if err != nil {
		panic(err)
	}
}

func BubbleSort(intSlice []int) []int {
	var ResultSlice []int = make([]int, len(intSlice))
	for len(intSlice) > 0 {
		var firstElement int
		var secondElement int
		for i := range intSlice {
			if i+1 >= len(intSlice) {
				continue
			}
			firstElement = intSlice[i]
			secondElement = intSlice[i+1]
			if intSlice[i] > intSlice[i+1] {
				intSlice[i] = secondElement
				intSlice[i+1] = firstElement
			}
		}

		ResultSlice[len(intSlice)-1] = intSlice[len(intSlice)-1]
		intSlice = intSlice[:len(intSlice)-1]
	}
	return ResultSlice
}
