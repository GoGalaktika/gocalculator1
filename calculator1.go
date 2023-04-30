package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	var (
		pattern    = `^(I|(II)|(III)|(IV)|V|(VI)|(VII)|(VIII)|(IX)|X)[\+\*\/\-](I|(II)|(III)|(IV)|V|(VI)|(VII)|(VIII)|(IX)|X)$|^(([1-9]|10)[\+\*\/\-]([1-9]|10))$`
		patternObj = regexp.MustCompile(pattern)
		reader     = bufio.NewReader(os.Stdin)
	)

	for {

		fmt.Println("Please enter a math expression or \"exit\" to quit (e.g. 5+4):")

		input, _ := reader.ReadString('\n')

		input = strings.ReplaceAll(input, " ", "")
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("Exiting program...")
			break
		} else if !(patternObj.MatchString(input)) {
			fmt.Println("Only +-*/ expressions of numbers from 1 to 10 are allowed")
			continue
		}

		opIndex := strings.IndexAny(input, "+-*/")

		op := input[opIndex]

		num1, _ := strconv.Atoi(input[0:opIndex])
		num2, _ := strconv.Atoi(input[opIndex+1:])

		//DO THE CALCULATIONS
		var result int
		switch op {
		case '+':
			result = num1 + num2
		case '-':
			result = num1 - num2
		case '*':
			result = num1 * num2
		case '/':
			result = num1 / num2
		default:
			fmt.Println("Invalid operator")
			continue
		}

		fmt.Printf("%v %c %v = %v\n", num1, op, num2, result) //print first number, operator and second number
		//
	}
}
