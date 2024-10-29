package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

// GrepFlags хранит флаги для утилиты grep
type GrepFlags struct {
	After      int
	Before     int
	Context    int
	Count      bool
	IgnoreCase bool
	Invert     bool
	Fixed      bool
	LineNum    bool
}

// grep выполняет поиск заданного шаблона в строках текста.
// Она возвращает строки, содержащие совпадения, а также количество найденных совпадений.
//
// Параметры:
//   - lines: срез строк, в которых будет осуществляться поиск.
//   - searchTerm: строка, которую необходимо найти.
//   - flags: структура, содержащая флаги поиска (например, для учета регистра,
//     игнорирования совпадений, вывода количества совпадений и т.д.)
//
// Возвращает:
// - срез строк с результатами поиска,
// - количество найденных совпадений,
// - ошибку (если произошла ошибка при компиляции регулярного выражения).
func grep(lines []string, searchTerm string, flags GrepFlags) ([]string, int, error) {
	var result []string
	var count int

	var re *regexp.Regexp
	var err error

	if flags.Fixed {
		re, err = regexp.Compile("^" + regexp.QuoteMeta(searchTerm) + "$")
	} else {
		if flags.IgnoreCase {
			re, err = regexp.Compile("(?i)" + searchTerm)
		} else {
			re, err = regexp.Compile(searchTerm)
		}
	}

	if err != nil {
		return nil, count, err
	}

	after := flags.After
	before := flags.Before
	if flags.Context > 0 {
		after = flags.Context
		before = flags.Context
	}

	for i, line := range lines {
		matched := re.MatchString(line)

		if flags.Invert {
			matched = !matched
		}

		if matched {
			if flags.Count {
				count++
			} else {
				start := max(0, i-before)
				end := min(len(lines)-1, i+after)

				for j := start; j <= end; j++ {
					lineOutput := lines[j]
					if flags.LineNum {
						lineOutput = fmt.Sprintf("%d: %s", j+1, lines[j])
					}
					result = append(result, lineOutput)
				}
			}
		}
	}

	return result, count, nil
}

func main() {
	flags := GrepFlags{}
	flag.IntVar(&flags.After, "A", 0, "Print N lines after the match")
	flag.IntVar(&flags.Before, "B", 0, "Print N lines before the match")
	flag.IntVar(&flags.Context, "C", 0, "Print N lines around the match")
	flag.BoolVar(&flags.Count, "c", false, "Print the number of matching lines")
	flag.BoolVar(&flags.IgnoreCase, "i", false, "Ignore case distinctions")
	flag.BoolVar(&flags.Invert, "v", false, "Select non-matching lines")
	flag.BoolVar(&flags.Fixed, "F", false, "Interpret searchTerm as a fixed string, not a regular expression")
	flag.BoolVar(&flags.LineNum, "n", false, "Show line number")
	flag.Parse()

	// Получение аргументов из командной строки
	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("Search term and file path are required")
		return
	}

	searchTerm := args[0]
	filepath := args[1]

	// Открытие файла для чтения
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		return
	}
	defer file.Close()

	// Чтение строк из файла
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Вызов функции grep для поиска
	matches, count, err := grep(lines, searchTerm, flags)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Вывод результатов поиска
	if flags.Count {
		fmt.Println(count)
	} else {
		for _, match := range matches {
			fmt.Println(match)
		}
	}
}

// min возвращает минимальное значение из двух целых чисел.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// max возвращает максимальное значение из двух целых чисел.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
