package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// sortString сортирует срез строк в лексикографическом порядке с использованием рекурсивного подхода.
func sortString(str []string) []string {
	if len(str) < 2 {
		return str
	}

	pivot := str[0]
	var left []string
	var right []string

	for i := 1; i < len(str); i++ {
		if str[i] < pivot {
			left = append(left, str[i])
		} else {
			right = append(right, str[i])
		}
	}

	return append(append(sortString(left), pivot), sortString(right)...)
}

// splitBySpaces разбивает строку на слова (колонки) по пробелам.
func splitBySpaces(s string) []string {
	return strings.Fields(s)
}

// convertStringsToInt конвертирует срез строк в срез целых чисел, возвращая ошибку, если конвертация не удалась.
func convertStringsToInt(str []string) ([]int, error) {
	var nums []int
	for _, s := range str {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}
	return nums, nil
}

// convertIntsToString конвертирует срез целых чисел в срез строк.
func convertIntsToString(nums []int) []string {
	var str []string
	for _, num := range nums {
		str = append(str, strconv.Itoa(num))
	}
	return str
}

// sortNumbers сортирует срез целых чисел в порядке возрастания с использованием рекурсивного подхода.
func sortNumbers(str []int) []int {
	if len(str) < 2 {
		return str
	}

	pivot := str[0]
	var left []int
	var right []int

	for i := 1; i < len(str); i++ {
		if str[i] < pivot {
			left = append(left, str[i])
		} else {
			right = append(right, str[i])
		}
	}

	return append(append(sortNumbers(left), pivot), sortNumbers(right)...)
}

// reverseString разворачивает порядок строк в срезе.
func reverseString(str []string) []string {
	result := make([]string, 0, len(str))
	for i := len(str) - 1; i >= 0; i-- {
		result = append(result, str[i])
	}
	return result
}

// sortUnique сортирует строки и удаляет дубликаты.
func sortUnique(str []string) []string {
	var unique []string
	uniqueMap := map[string]struct{}{}

	for _, item := range str {
		if _, exist := uniqueMap[item]; !exist {
			unique = append(unique, item)
			uniqueMap[item] = struct{}{}
		}
	}

	return unique
}

// sortMonths сортирует строки, представляющие месяцы, в порядке их следования в календаре.
func sortMonths(months []string) []string {
	var monthOrder = map[string]int{
		"January":   1,
		"February":  2,
		"March":     3,
		"April":     4,
		"May":       5,
		"June":      6,
		"July":      7,
		"August":    8,
		"September": 9,
		"October":   10,
		"November":  11,
		"December":  12,
	}

	if len(months) < 2 {
		return months
	}

	pivot := months[0]
	var left, right []string

	for i := 1; i < len(months); i++ {
		if monthOrder[months[i]] < monthOrder[pivot] {
			left = append(left, months[i])
		} else {
			right = append(right, months[i])
		}
	}

	return append(append(sortMonths(left), pivot), sortMonths(right)...)
}

// trimRightSpace удаляет пробелы в конце каждой строки в срезе.
func trimRightSpace(str []string) []string {
	trimmedSlice := make([]string, len(str))

	for i, s := range str {
		trimmedSlice[i] = strings.TrimRight(s, " ")
	}
	return trimmedSlice
}

func isSorted(str []string) bool {
	for i := 1; i < len(str); i++ {
		if str[i] < str[i-1] {
			return false
		}
	}
	return true
}

// parseNumberWithSuffix разбирает числовые значения с суффиксами (K, M, B) и возвращает их в виде float64.
func parseNumberWithSuffix(s string) (float64, error) {
	suffixes := map[rune]float64{
		'K': 1000,
		'M': 1_000_000,
		'B': 1_000_000_000,
	}
	s = strings.TrimSpace(s)
	numStr := ""
	suffix := ' '
	for _, r := range s {
		if unicode.IsDigit(r) || r == '.' {
			numStr += string(r)
		} else {
			suffix = r
			break
		}
	}
	num, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		return 0, err
	}
	if scale, ok := suffixes[suffix]; ok {
		return num * scale, nil
	}
	return num, nil
}

// sortNumbersWithSuffix сортирует строки с числовыми значениями, учитывая суффиксы.
func sortNumbersWithSuffix(str []string) ([]string, error) {
	var nums []float64
	for _, s := range str {
		num, err := parseNumberWithSuffix(s)
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}

	// Сортировка чисел
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j] // Меняем местами числа
				str[j], str[j+1] = str[j+1], str[j]     // Меняем местами строки
			}
		}
	}

	return str, nil
}

// sortLinesByKey сортирует строки на основе переданных ключей (параметров).
// Поддерживаемые ключи:
// -k (по указанной колонке),
// -n (по числам),
// -r (реверс),
// -u (уникальные),
// -M (по месяцам),
// -b (обрезка пробелов),
// -c (проверять отсортированы ли данные),
// -h (по числам с суффиксами).
func sortLinesByKey(str []string, keys ...string) []string {
	if len(keys) == 0 {
		return sortString(str)
	}

	result := str

	for _, key := range keys {
		switch key {
		case "-k":
			if len(keys) > 1 {
				colIndex, err := strconv.Atoi(keys[1])
				if err == nil && colIndex > 0 {
					colIndex--
					var columnValues []string
					for _, line := range result {
						cols := splitBySpaces(line)
						if colIndex < len(cols) {
							columnValues = append(columnValues, cols[colIndex])
						}
					}
					result = sortString(columnValues)
				}
			}
		case "-n":
			nums, err := convertStringsToInt(result)
			if err != nil {
				fmt.Println(err)
				return nil
			}
			resultNums := sortNumbers(nums)
			result = convertIntsToString(resultNums)
		case "-r":
			if len(keys) == 1 {
				result = reverseString(sortString(result))
			} else {
				result = reverseString(result)
			}
		case "-u":
			result = sortUnique(result)
		case "-M":
			result = sortMonths(result)
		case "-b":
			result = trimRightSpace(result)
		case "-c":
			if isSorted(result) {
				fmt.Println("The lines are sorted.")
			} else {
				fmt.Println("The lines are NOT sorted.")
			}

		case "-h":
			sorted, err := sortNumbersWithSuffix(result)
			if err != nil {
				fmt.Println(err)
				return nil
			}
			result = sorted

		}
	}

	return result
}

func main() {
	var line []string
	var result []string

	line = []string{"banana", "apple", "orange"}
	result = sortLinesByKey(line)
	fmt.Println(result)

	line = []string{"apple banana", "orange apple", "banana orange"}
	result = sortLinesByKey(line, "-k", "2")
	fmt.Println(result)

	line = []string{"5", "2", "13"}
	result = sortLinesByKey(line, "-n")
	fmt.Println(result)

	line = []string{"5", "2", "13", "2", "10"}
	result = sortLinesByKey(line, "-r")
	fmt.Println(result)

	line = []string{"5", "2", "13", "2", "10"}
	result = sortLinesByKey(line, "-u")
	fmt.Println(result)

	line = []string{"March", "January", "February", "December", "July", "April"}
	result = sortLinesByKey(line, "-M")
	fmt.Println(result)

	line = []string{"   leading space", "trailing space   ", "  both  "}
	result = sortLinesByKey(line, "-b")
	fmt.Println(result)

	line = []string{"5K", "2M", "13B", "2", "10"}
	result = sortLinesByKey(line, "-h")
	fmt.Println(result)

	line = []string{"banana", "apple", "orange"}
	result = sortLinesByKey(line, "-c")
	fmt.Println(result)
}
