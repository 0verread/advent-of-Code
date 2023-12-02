package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func sumOfFirstLastDigit(str string) string {

	//iterate over string to create a new string
	// if current substring has any digit in words, replace with digit
	// now make operation

	firstD := ' '
	lastD := ' '
	// tempLastD := ' '
	digitsInWords := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	subStrPos := -1
	m := map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}

	var newStr string
	for _, c := range str {
		newStr += string(c)
		// fmt.Println(newStr)
		if len(newStr) > 2 {
			for _, value := range digitsInWords {
				subStrPos = strings.Index(newStr, value)
				if subStrPos != -1 {
					newStr = strings.Replace(newStr, value, m[value], 1)
					newStr += value[len(value)-1:]
				}
			}
		}
	}

	for _, c := range newStr {
		// fmt.Println(string(firstD), string(tempLastD))
		fmt.Println(str, " ", newStr)
		if unicode.IsDigit(c) {
			if firstD == ' ' {
				firstD = c
			}
			lastD = c
		}
	}

	return string(firstD) + string(lastD)
}

func main() {
	file, err := os.Open("./input/day1.txt")
	// file, err := os.Open("./input/day1.test.txt")
	if err != nil {
		fmt.Println("Some error in opening file")
	}

	defer file.Close()

	var res int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// dont add if return sum is negative
		rNum := sumOfFirstLastDigit(scanner.Text())
		if rNum != " " {
			iNum, _ := strconv.Atoi(rNum)
			fmt.Println(iNum)
			res += iNum
		}

	}
	fmt.Println(res)
}
