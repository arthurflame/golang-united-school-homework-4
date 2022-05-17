package string_sum

import (
	"errors"
	"fmt"
	"strconv"
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

	numbers, err := getNumbers(output)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	converted := convertNumbers(numbers)

	output = strconv.Itoa(converted)

	return output, nil
}

func cleanInput(input string) (string, error) {
	var result string

	for i := range input {

		if string(input[i]) == "+" || string(input[i]) == "-" || unicode.IsDigit(rune(input[i])) {
			result += string(input[i])
		} else {
			return "", fmt.Errorf("char: %q, can't be parsed", string(input[i]))

		}
	}
	result = strings.Join(strings.Split(strings.TrimSpace(input), " "), "")
	return result, nil

}

func splitAny(s string, sep string) []string {
	splitter := func(r rune) bool {
		return strings.ContainsRune(sep, r)
	}
	return strings.FieldsFunc(s, splitter)
}

func getNumbers(input string) ([]string, error) {
	result := make([]string, 0)
	var tmp string

	for i := range input {
		if string(input[i]) == "-" {
			tmp += "." + string(input[i])
		} else {
			tmp += string(input[i])
		}
	}

	result = splitAny(tmp, ".+")

	if len(result) > 2 || len(result) <= 1 {
		return result, fmt.Errorf("%w", errorNotTwoOperands)
	}

	return result, nil
}

func convertNumbers(input []string) int {
	var res int
	for i := range input {
		if s, err := strconv.Atoi(input[i]); err == nil {
			res += s
		}
	}

	return res
}
