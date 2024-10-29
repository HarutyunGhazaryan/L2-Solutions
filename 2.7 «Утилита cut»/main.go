package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// GrepFlags хранит флаги для утилиты grep
type CutFlags struct {
	Fields    string
	Delimiter string
	Separated bool
	FilePath  string
}

// getSelectedFields принимает срез строк `columns` и структуру флагов `flags`,
// извлекает выбранные поля из `columns` на основе индексов, указанных в `flags.Fields`.
// Функция проверяет корректность индексов и возвращает ошибку, если индекс недопустим (например,
// выходит за пределы доступных столбцов). Если индексы действительны, функция возвращает срез
// строк, содержащий выбранные поля.
func getSelectedFields(columns []string, flags CutFlags) ([]string, error) {
	var result []string

	if flags.Fields != "" {
		fields := strings.Split(flags.Fields, ",")
		for _, value := range fields {
			index, err := strconv.Atoi(value)
			if err != nil {
				return nil, fmt.Errorf("invalid field index: %s", value)
			}
			if index < 1 || index > len(columns) {
				return nil, fmt.Errorf("field index %d out of range", index)
			}
			result = append(result, columns[index-1])
		}
	}
	return result, nil
}

func main() {
	flags := CutFlags{}
	flag.StringVar(&flags.Fields, "f", "", "Select fields")
	flag.StringVar(&flags.Delimiter, "d", "\t", "Use a different separator")
	flag.BoolVar(&flags.Separated, "s", false, "Only lines with separator")
	flag.StringVar(&flags.FilePath, "file", "", "Path to the text file")
	flag.Parse()

	var scanner *bufio.Scanner

	// Открываем файл, если путь к нему передан, иначе читаем из STDIN
	if flags.FilePath != "" {
		file, err := os.Open(flags.FilePath)
		if err != nil {
			fmt.Printf("Error opening file: %s\n", err)
			return
		}

		defer file.Close()
		scanner = bufio.NewScanner(file)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}

	// Чтение строк из файла
	for scanner.Scan() {
		line := scanner.Text()

		if flags.Separated && !strings.Contains(line, flags.Delimiter) {
			continue
		}

		columns := strings.Split(line, flags.Delimiter)

		// Вызов функции getSelectedFields для поиска
		selectedFields, err := getSelectedFields(columns, flags)
		if err != nil {
			fmt.Printf("Error processing line: %s\n", err)
			continue
		}
		fmt.Println(strings.Join(selectedFields, flags.Delimiter))

	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input: %s\n", err)
	}

}
