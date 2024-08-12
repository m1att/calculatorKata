package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func sum(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			panic("На ноль делить нельзя")
		}
		return a / b
	default:
		panic("Неизвестный оператор")
	}
}

func isValidRoman(roman string) bool {
	
	validRomanPattern := `^M{0,3}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$`
	matched, _ := regexp.MatchString(validRomanPattern, roman)
	return matched
}

func romanToArabic(roman string) int {
	if !isValidRoman(roman) {
		panic("Неверный формат римского числа")
	}

	elements := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100}

	total := 0
	prev := 0
	for _, ch := range roman {
		curr := elements[ch]
		if curr > prev {
			total += curr - 2*prev
		} else {
			total += curr
		}
		prev = curr
	}
	return total
}

func arabicToRoman(arabic int) string {
	elements := []struct {
		Value  int
		Symbol string
	}{
		{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
		{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
	}

	var result strings.Builder
	for _, element := range elements {
		for arabic >= element.Value {
			result.WriteString(element.Symbol)
			arabic -= element.Value
		}
	}
	return result.String()
}

func checkType(value string) (int, bool) {
	
	if arabic, err := strconv.Atoi(value); err == nil && arabic >= 1 && arabic <= 10 {
		return arabic, false
	}
	
	if !isValidRoman(value) {
		panic("Неверный формат римского числа")
	}
	
	return romanToArabic(value), true
}

func main() {
	fmt.Println("Введите пример (например, 10 + 5 или X + V):")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	parts := strings.Fields(input)
	if len(parts) != 3 {
		fmt.Println("Неверный формат примера.")
		return
	}

	firstValue, isRoman1 := checkType(parts[0])
	secondValue, isRoman2 := checkType(parts[2])
	operator := parts[1]

	if isRoman1 != isRoman2 {
		panic("Используйте только арабские или только римские цифры.")
	}

	result := sum(firstValue, secondValue, operator)

	if isRoman1 {
		if result <= 0 {
			panic("Результат не может быть отрицательным или нулем для римских цифр.")
		}
		fmt.Println(arabicToRoman(result))
	} else {
		fmt.Println(result)
	}
}
