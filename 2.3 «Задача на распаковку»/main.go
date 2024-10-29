package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// unpackStringWithBuilder распаковывает строку, используя strings.Builder.
// Функция обрабатывает символы и руны в строке. Если встречается
// буква, она сохраняется. Если встречается число, предыдущая буква
// повторяется соответствующее количество раз. Поддерживаются
// escape-последовательности для букв, а также некорректные форматы
// строк обрабатываются через возврат ошибки.
func unpackStringWithBuilder(s string) (string, error) {
	var result strings.Builder
	var lastLetter string
	runes := []rune(s)
	escaped := false

	for i := 0; i < len(runes); i++ {
		char := runes[i]

		if char == '\\' {
			if escaped {
				result.WriteString("\\")
				lastLetter = "\\"
				escaped = false
			} else {
				escaped = true
			}
			continue
		}

		if escaped {
			result.WriteString(string(char))
			lastLetter = string(char)
			escaped = false
			continue
		}

		if unicode.IsLetter(char) {
			result.WriteString(string(char))
			lastLetter = string(char)
		} else if unicode.IsDigit(char) {
			if lastLetter == "" {
				return "", errors.New("incorrect string format")
			}
			num, err := strconv.Atoi(string(char))
			if err != nil {
				return "", err
			}

			for j := 0; j < num-1; j++ {
				result.WriteString(lastLetter)
			}
		} else {
			return "", errors.New("incorrect string format")
		}
	}

	return result.String(), nil
}

// unpackStringWithConcat распаковывает строку, используя конкатенацию строк.
// Функция работает аналогично unpackStringWithBuilder, но использует
// конкатенацию строк для формирования результирующей строки. Она обрабатывает
// буквы и цифры, повторяя предыдущую букву необходимое количество раз. Также
// поддерживаются escape-последовательности, и некорректные форматы строк
// возвращают ошибку.
func unpackStringWithConcat(s string) (string, error) {
	var result string
	var lastLetter string
	runes := []rune(s)
	escaped := false

	for i := 0; i < len(runes); i++ {
		char := runes[i]

		if char == '\\' {
			if escaped {
				result += "\\"
				lastLetter = "\\"
				escaped = false
			} else {
				escaped = true
			}
			continue
		}

		if escaped {
			result += string(char)
			lastLetter = string(char)
			escaped = false
			continue
		}

		if unicode.IsLetter(char) {
			result += string(char)
			lastLetter = string(char)
		} else if unicode.IsDigit(char) {
			if lastLetter == "" {
				return "", errors.New("incorrect string format")
			}
			num, err := strconv.Atoi(string(char))
			if err != nil {
				return "", err
			}

			for j := 0; j < num-1; j++ {
				result += lastLetter
			}
		} else {
			return "", errors.New("incorrect string format")
		}
	}

	return result, nil
}

func main() {
	testStrings := []string{
		"a4bc2d5e",
		"abcd",
		"45",
		"",
		"qwe\\4\\5",
		"qwe\\45",
		"qwe\\\\5",
	}

	for _, s := range testStrings {
		unpacked, err := unpackStringWithBuilder(s)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println(unpacked)
		}

		unpacked, err = unpackStringWithConcat(s)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println(unpacked)
		}
	}
}
