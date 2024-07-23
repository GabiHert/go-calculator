package operations

func Subtract[T float64 | float32 | int | rune | complex64](n ...T) T {
	var res T
	//Obs: for index, element := range array
	for _, f := range n {
		res -= f
	}
	return res
}
