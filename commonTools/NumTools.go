package commonTools

func GetMinimum(intSlice []int) int {
	var minimum int = int(^uint(0) >> 1)

	for _, item := range intSlice {
		if minimum > item {
			minimum = item
		}
	}

	return minimum
}

func MakeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func ChooseSmaller(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
