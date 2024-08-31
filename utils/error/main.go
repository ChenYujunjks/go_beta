package eroeoeo

import (
	"errors"
	"fmt"
)

func maxInArray(arr []int) (int, error) {
	if len(arr) == 0 {
		return 0, errors.New("array is empty")
	}

	m := arr[0]
	for _, value := range arr {
		if value > m {
			m = value
		}
	}
	return m, nil
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func Error_F() {
	arr := []int{}
	max, err := maxInArray(arr)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("The maximum value is:", max)
	}

	result, err := divide(4, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	result, err = divide(4, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
	name := "world"
	er2 := fmt.Errorf("hello, %s", name)
	fmt.Println(er2)
	fmt.Println("===============================")
}
