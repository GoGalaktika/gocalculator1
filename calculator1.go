package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func convertRomanToArab(romNum string) int {
	var arabNum int
	for i := 0; i < len(romNum); i++ {
		for j := 0; j < len(romanMap); j++ {
			if romNum[i] == romanMap[j] {
				return
			}
		}
	}
}

func main() {

	var (
		pattern            = `^(I|(II)|(III)|(IV)|V|(VI)|(VII)|(VIII)|(IX)|X)[\+\*\/\-](I|(II)|(III)|(IV)|V|(VI)|(VII)|(VIII)|(IX)|X)$|^(([1-9]|10)[\+\*\/\-]([1-9]|10))$`
		patternObj         = regexp.MustCompile(pattern)
		reader             = bufio.NewReader(os.Stdin)
		romanMap           = map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}
		input              string
		op                 byte
		convNum1, convNum2 int
	)

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
		}

		convNum1, _ = strconv.Atoi(input[0:opIndex])
		convNum2, _ = strconv.Atoi(input[opIndex+1:])

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
