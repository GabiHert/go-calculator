package operations

import "errors"

func Divide[T float64 | float32 | int | rune | complex64](a, b T) (res T, err error) {
	//Obs: 'T' is our generic type, it's a type constraint that matches float32, float64, int, and complex64
	//Obs: 'res' and 'err' are named returns, so they aren't explicitly initialized
	if b == 0 {
		err = errors.New("can't divide by zero")
		return //Obs: same as "return err, res" since 'res' and 'err' are named returns
	}

	res = a / b
	return res, nil
}
