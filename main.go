package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Read() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Что посчитать?")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	text = strings.ReplaceAll(text, " ", "")
	return text
}

func Preform(text string) (string, string, string) {
	f := func(c rune) bool {
		if c == '+' || c == '-' || c == '*' || c == '/' {
			return true
		}
		return false
	}
	i := strings.IndexFunc(text, f)
	if i == -1 {
		panic("Error: Unknown operator or not a mathematical operation")
	}
	if strings.ContainsFunc(text[i+1:], f) {
		panic("Error: Wrong format of operation")
	}
	return text[:i], text[i+1:], string(text[i])
}

func Make(a, b string, op string) interface{} {
	var res int
	x, err := strconv.Atoi(a)
	y, err2 := strconv.Atoi(b)
	if err == nil && err2 == nil {
		if x <= 0 || y <= 0 || x > 10 || y > 10 {
			panic("Error: Some numbers are > 10 or = 0")
		}
		res = calculateNums(x, y, op)
	} else if err != nil && err2 != nil {
		x = Transform(a)
		y = Transform(b)
		if op == "-" && y > x {
			panic("Error: Negative result in the Roman system ")
		}
		res = calculateNums(x, y, op)
		return fromNumsToRoman(res)
	} else {
		panic("Error: Different types of nums")
	}
	return res
}

func calculateNums(a int, b int, sign string) int {
	switch sign {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	}
	return 0
}

func Transform(a string) int {
	var x int
	switch a {
	case "I":
		x = 1
	case "II":
		x = 2
	case "III":
		x = 3
	case "IV":
		x = 4
	case "V":
		x = 5
	case "VI":
		x = 6
	case "VII":
		x = 7
	case "VIII":
		x = 8
	case "IX":
		x = 9
	case "X":
		x = 10
	default:
		panic("Error: Unknown number")
	}
	return x
}

func fromNumsToRoman(num int) string {
	var x string

	for num > 0 {
		switch {
		case num-100 == 0:
			x += "C"
			num -= 100
		case num-90 >= 0:
			x += "XC"
			num -= 90
		case num-50 >= 0:
			x += "L"
			num -= 50
		case num-40 >= 0:
			x += "XL"
			num -= 40
		case num-10 >= 0:
			x += "X"
			num -= 10
		case num-9 >= 0:
			x += "IX"
			num -= 9
		case num-5 >= 0:
			x += "V"
			num -= 5
		case num-4 >= 0:
			x += "IV"
			num -= 4
		case num-1 >= 0:
			x += "I"
			num -= 1
		default:
			panic("Error: Unknown number")
		}
	}
	return x
}

func main() {
	text := Read()
	x, y, op := Preform(text)
	res := Make(x, y, op)
	fmt.Println(res)
}
