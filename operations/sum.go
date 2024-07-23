package operations

func Sum(n ...float64) float64 {
	var res float64
	for _, f := range n {
		res += f
	}
	return res

}
