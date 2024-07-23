package operations

import "fmt"

func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("cant calculate a factorial of a negative number")
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result, nil
}
