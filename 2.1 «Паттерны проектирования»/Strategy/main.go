package main

import (
	"fmt"
	"sort"
)

// SortStrategy интерфейс для стратегий сортировки
type SortStrategy interface {
	SortNumbers([]int) []int
}

type AscendingSortStrategy struct{}

func (s AscendingSortStrategy) SortNumbers(numbers []int) []int {
	sorted := make([]int, len(numbers))
	copy(sorted, numbers)
	sort.Ints(sorted)
	return sorted
}

type DescendingSortStrategy struct{}

type ByDescending []int

func (a ByDescending) Len() int           { return len(a) }
func (a ByDescending) Less(i, j int) bool { return a[i] > a[j] }
func (a ByDescending) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func (s DescendingSortStrategy) SortNumbers(numbers []int) []int {
	sorted := make([]int, len(numbers))
	copy(sorted, numbers)
	sort.Sort(ByDescending(sorted))
	return sorted
}

// Numbers структура, содержащая числа и стратегию сортировки
type NumberCollection struct {
	Numbers  []int
	Strategy SortStrategy
}

// NewSort функция для создания нового объекта Numbers с заданной стратегией
func NewNumberCollection(numbers []int, strategy SortStrategy) *NumberCollection {
	return &NumberCollection{
		Numbers:  numbers,
		Strategy: strategy,
	}
}

func (n NumberCollection) SortNumbers() []int {
	return n.Strategy.SortNumbers(n.Numbers)
}

func main() {
	numbers := []int{1, 3, 10, 5, 7}

	ascendingSortStrategy := &AscendingSortStrategy{}
	descendingSortStrategy := &DescendingSortStrategy{}

	numberCollection1 := NewNumberCollection(numbers, ascendingSortStrategy)
	fmt.Println("Sorting in ascending order: ", numberCollection1.SortNumbers())

	numberCollection2 := NewNumberCollection(numbers, descendingSortStrategy)
	fmt.Println("Sorting in descending order", numberCollection2.SortNumbers())
}

// Паттерн "Стратегия" позволяет изменять поведение объекта во время выполнения программы,
// подставляя в него различные алгоритмы или стратегии. Это делается через определение
// интерфейса для алгоритмов и создание конкретных реализаций, которые могут быть
// использованы в зависимости от ситуации.

// Плюсы применения паттерна "Стратегия":
// 1. Позволяет варьировать поведение объекта.
//    Стратегия позволяет динамически менять алгоритм, используемый объектом,
//    в зависимости от условий или входных данных.

// 2. Изолирует код алгоритмов от других классов.
//    Код, отвечающий за алгоритмы, отделяется от остального кода, что позволяет
//    скрыть детали реализации и улучшить читаемость программы.

// 3. Упрощает добавление новых алгоритмов.
//    Для добавления нового поведения достаточно создать новую реализацию интерфейса
//    стратегии, не изменяя основной код программы. Это соответствует принципу
//    открытости/закрытости.

// Минусы применения паттерна "Стратегия":
// 1. Усложняет программу за счёт дополнительных структур.
//    Каждая стратегия требует создания своей структуры, что может привести к
//    увеличению количества структур и усложнению архитектуры программы.

// 2. Клиент должен знать различия между стратегиями.
//    Для правильного выбора стратегии клиентский код должен понимать, как
//    разные стратегии влияют на поведение объекта, что может усложнить использование
//    паттерна в некоторых ситуациях.
