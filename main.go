package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

var maxAllowedNum = GetMax()
var ErrHighValue = errors.New("too high value error")
var ErrNegativeValue = errors.New("negative value error")

func main() {
	testValues := [][]string{
		{"-1", "Значение не может быть отрицательным"},
		{"0", "0"},
		{"1", "0"},
		{"2", "2"},
		{"4", "6"},
		{"85672", "1834965732"},
		{"92683", "2147534622"},
		{"92684", "2147627306"},
		{"92685", "Значение не может быть больше " + strconv.Itoa(maxAllowedNum)},
		{"101248", "Значение не может быть больше " + strconv.Itoa(maxAllowedNum)},
		{"one", "Значение не является числом"},
		{"СЕМЬ", "Значение не является числом"},
		{"~`!@#$%^&*()_+-=][}{\\|\"':;/?<>.,", "Значение не является числом"},
	}

	for i := 0; i < len(testValues); i++ {
		testValue, err := strconv.Atoi(testValues[i][0])
		expectedValue := testValues[i][1]
		if err != nil {
			fmt.Printf("Test value: %s; Expected value: %s; Result value: Значение не является числом\n", testValues[i][0], expectedValue)
			continue
		}

		result, err := EvenNumbersSum(testValue)
		if err != nil {
			switch {
			case errors.Is(err, ErrNegativeValue):
				fmt.Printf("Test value: %d; Expected value: %s; Result value: Значение не может быть отрицательным\n", testValue, expectedValue)
			case errors.Is(err, ErrHighValue):
				fmt.Printf("Test value: %d; Expected value: %s; Result value: Значение не может быть больше %d\n", testValue, expectedValue, maxAllowedNum)
			default:
				fmt.Println("Что-то пошло не так")
			}
		} else {
			fmt.Printf("Test value: %d; Expected value: %s; Result value: %d\n", testValue, expectedValue, result)
		}
	}
}

func EvenNumbersSum(num int) (int, error) {
	result := 0

	if (num >= 0) && (num <= maxAllowedNum) {
		for i := 2; i <= num; i += 2 {
			result += i
		}
	} else if num < 0 {
		return 0, ErrNegativeValue
	} else {
		return 0, ErrHighValue
	}

	return result, nil
}

func GetMax() int {
	result := 0
	i := 2

	for result <= math.MaxInt32 {
		result += i
		i += 2
	}

	return i
}
