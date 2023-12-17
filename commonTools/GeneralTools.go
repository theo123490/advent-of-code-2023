package commonTools

func PrintErr(err error) {
	if err != nil {
		panic(err)
	}
}
