package main

import "fmt"

type InfoLogger interface {
	LogInfo(info string)
}

type ErrorLogger interface {
	logErr(err string)
}

type LogInformation struct{}

func (l LogInformation) LogInfo(info string) {
	fmt.Println(info)
}

type LogError struct{}

func (l LogError) logErr(err string) {
	fmt.Println(err)
}

// Общий интерфейс фабрики для создания логгеров
type LoggerFactory interface {
	CreateInfoLogger() InfoLogger
	CreateErrorLogger() ErrorLogger
}

// Конкретная фабрика для логгеров
type ConcreteLoggerFactory struct{}

func (f ConcreteLoggerFactory) CreateInfoLogger() InfoLogger {
	return LogInformation{}
}

func (f ConcreteLoggerFactory) CreateErrorLogger() ErrorLogger {
	return LogError{}
}

func main() {
	factory := ConcreteLoggerFactory{}

	infoLogger := factory.CreateInfoLogger()
	infoLogger.LogInfo("Application started successfully")

	errorLogger := factory.CreateErrorLogger()
	errorLogger.logErr("Failed to connect to the database")
}

// Паттерн "Фабричный метод" позволяет отделить код создания объектов (структур)
// от кода, который эти объекты использует. Это делается через определение
// интерфейса или метода для создания продуктов (структур), что позволяет
// различным типам продуктов реализовывать этот метод по-своему.

// Плюсы применения паттерна "Фабричный метод":
// 1. Избавляет структуры от жёсткой привязки к конкретным типам продуктов.
//    Основной код не зависит от конкретных типов создаваемых структур.
//    Это упрощает изменение и расширение программы, так как можно добавлять
//    новые типы продуктов (структур), не изменяя основной код.

// 2. Выделяет код создания продуктов в одно место.
//    Логика создания объектов сосредотачивается в фабричных функциях или методах,
//    что делает код более структурированным и поддерживаемым.

// 3. Упрощает добавление новых продуктов (структур).
//    Для добавления нового типа продукта достаточно создать новую реализацию
//    интерфейса или метод создания, не изменяя основной код программы.
//    Это соответствует принципу открытости/закрытости.

// Минусы применения паттерна "Фабричный метод":
// 1. Может привести к усложнению структуры программы.
//    Для каждого типа продукта нужно создавать свои структуры и фабрики,
//    что может увеличить количество кода и усложнить иерархию программы.

// 2. Может потребовать больше кода для добавления новых типов продуктов.
//    Для каждого нового типа продукта нужно реализовать новые структуры и фабричные методы,
//    что может привести к увеличению количества кода при расширении системы.
