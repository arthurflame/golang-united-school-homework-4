package string_sum

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

//if strings.Contains(input, " ") {
//	return output, fmt.Errorf("%w", errorEmptyInput)
//}

func StringSum(input string) (output string, err error) {
	if input == " " {
		return input, fmt.Errorf("%w", errorEmptyInput)
	}

	output, err = cleanInput(input)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}
	fmt.Println("cleaned output:", output)

	z := getNumbers(output)
	fmt.Println(z)

	//getOperands()

	return "", nil
}

func cleanInput(input string) (string, error) {
	var result string

	for i := range input {
		// -2+3
		if string(input[i]) == "+" || string(input[i]) == "-" || unicode.IsDigit(rune(input[i])) {
			result += string(input[i])
		} else {
			return "", fmt.Errorf("char: %q, can't be parsed", string(input[i]))

		}
	}
	result = strings.Join(strings.Split(strings.TrimSpace(input), " "), "")
	return result, nil

}

func getNumbers(input string) []int {
	result := make([]string, 0)
	var temp string

	//if strings.HasPrefix(input, "-") {
	//	i = 1
	//}

	for i := range input {
		if strings.Count(input, "-") < 2 && string(input[i]) != "+" {
			temp += string(input[i])
		} else {
			result = append(result, temp)
		}
		if string(input[i]) == "+" && strings.Count(strings.Join(result, ""), "-") > 1 {
			temp += string(input[i])
		}
		result = append(result, temp)
	}

	fmt.Println(result)
	return []int{0}
}

//func getOperands(input string) []int {
//	return []int{0}
//}
