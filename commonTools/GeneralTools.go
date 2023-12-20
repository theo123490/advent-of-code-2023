package commonTools

func PrintErr(err error) {
	if err != nil {
		panic(err)
	}
}

type Valuables interface {
	IsHigher(Valuables) bool
}

func BubbleSort(valueSlice []Valuables) []Valuables {
	var resultSlice []Valuables = make([]Valuables, len(valueSlice))
	for len(valueSlice) > 0 {
		var firstElement Valuables
		var secondElement Valuables
		for i := range valueSlice {
			if i+1 >= len(valueSlice) {
				continue
			}
			firstElement = valueSlice[i]
			secondElement = valueSlice[i+1]
			if valueSlice[i].IsHigher(valueSlice[i+1]) {
				valueSlice[i] = secondElement
				valueSlice[i+1] = firstElement
			}
		}

		resultSlice[len(valueSlice)-1] = valueSlice[len(valueSlice)-1]
		valueSlice = valueSlice[:len(valueSlice)-1]
	}
	return resultSlice
}
