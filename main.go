package main

import (
	"calculator/operations"
	"fmt"
)

func main() {
	res := operations.SafeSum(0.1, 0.2)
	fmt.Printf("Result %s \n", res.StringFixed(20))
}
