package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	pattern            = `^(I|(II)|(III)|(IV)|V|(VI)|(VII)|(VIII)|(IX)|X)[\+\*\/\-](I|(II)|(III)|(IV)|V|(VI)|(VII)|(VIII)|(IX)|X)$|^(([1-9]|10)[\+\*\/\-]([1-9]|10))$`
	patternObj         = regexp.MustCompile(pattern)
	reader             = bufio.NewReader(os.Stdin)
	romanArabMap       = map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}
	input              string
	op                 byte
	convNum1, convNum2 int
)

func convertRomanToArab(romNum string) int {
	for key, value := range romanArabMap {
		if romNum == key {
			return value
		}
	}
	return 0
}

func convertArabToRoman(arabNum int) string {
	for key, value := range romanArabMap {
		if arabNum == value {
			return key
		}
	}
	return ""
}

func main() {

	for {

		fmt.Println("Please enter a math expression or \"exit\" to quit (e.g. 5+4 or IV*III):")

		input, _ = reader.ReadString('\n')

		input = strings.ReplaceAll(input, " ", "")
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("Exiting program...")
			break
		} else if !(patternObj.MatchString(input)) {
			fmt.Println("Only +-*/ expressions of roman and arabic numbers from 1 to 10 are allowed")
			continue
		}

		opIndex := strings.IndexAny(input, "+-*/")

		op = input[opIndex]

		num1 := input[0:opIndex]
		num2 := input[opIndex+1:]

		if !(regexp.MustCompile(`^\d+$`).MatchString(num1)) {
			convNum1 = convertRomanToArab(num1)
			convNum2 = convertRomanToArab(num2)
		} else {
			convNum1, _ = strconv.Atoi(input[0:opIndex])
			convNum2, _ = strconv.Atoi(input[opIndex+1:])
		}

		//DO THE CALCULATIONS
		var result int
		switch op {
		case '+':
			result = convNum1 + convNum2
		case '-':
			result = convNum1 - convNum2
		case '*':
			result = convNum1 * convNum2
		case '/':
			result = convNum1 / convNum2
		default:
			fmt.Println("Invalid operator")
			continue
		}

		fmt.Printf("%v %T %c %v %T = %v\n", num1, num1, op, num2, num2, result) //print first number, operator and second number
		//
	}
}
