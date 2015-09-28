package main

import ("fmt"; "errors")


var InvalidSumError = errors.New("Invalid Sum, only positive values are required")

func PositiveSum(i, y int) (int, error) {
	defer func() {
		fmt.Printf("Call after function return\n")
	}()

	if i < 0 || y < 0 {
		return 0, InvalidSumError
	}

	return i + y, nil
}

func main() {
	if value, err := PositiveSum(1, -1); InvalidSumError == err {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Value: %+v\n", value)
	}

	value, err := PositiveSum(1, 1)

	fmt.Printf("Value: %+v, Error: %s\n", value, err)
}