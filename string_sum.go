package string_sum

import (
	"errors"
	"fmt"
	"strings"
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

	//var operands []string

	output = strings.Join(strings.Split(strings.TrimSpace(input), " "), "")
	//fmt.Println(output)
	getOperands(output)

	return "", nil
}

// input := "-2+5"
// -,+,-
// 2,3,4
// isFirstOpend bool  input[0] == "-"
//result:=

func getOperands(input string) []int {
	var operators = make([]string, len(input))
	var operands = make([]string, len(input))

	if string(input[0]) == "-" {
		operators[0] = "-"
		for i := 1; i < len(input); i++ {
			if string(input[i]) == "-" || string(input[i]) == "+" {
				operators[i] = string(input[i])
			}
		}
	} else {
		for i := 0; i < len(input); i++ {
			if string(input[i]) == "-" || string(input[i]) == "+" {
				operators[i] = string(input[i])
			}
		}
	}

	if strings.HasPrefix(input, "-") {
		for i := 1; i < len(input); i++ {
			if string(input[i]) != "-" && string(input[i]) != "+" {
				operands[i] = string(input[i])
			}
		}
	} else {
		for i := 0; i < len(input); i++ {
			if string(input[i]) != "-" && string(input[i]) != "+" {
				operands[i] = string(input[i])
			}
		}
	}

	fmt.Println("operators:", operators)
	fmt.Println("operands:", operands)
	
	return []int{0}
}
