package main

import (
	"fmt"
	"strings"
)

// Определение типа ошибки валидации
type ValidationError struct {
	Message string
}

// Реализация метода Error для ошибки валидации
func (v *ValidationError) Error() string {
	return v.Message
}

// Структура формы регистрации
type Registration struct {
	email    string
	password string
}

// Интерфейс для обработчика
type Handler interface {
	SetNext(handler Handler) Handler // Устанавливает следующий обработчик
	Handle(form *Registration) error // Выполняет обработку формы
}

// Базовый обработчик с полем для следующего в цепочке
type BaseHandle struct {
	next Handler
}

// Метод установки следующего обработчика
func (h *BaseHandle) SetNext(handler Handler) Handler {
	h.next = handler
	return handler
}

// Метод обработки, передает выполнение следующему обработчику, если он существует
func (h *BaseHandle) Handle(form *Registration) error {
	if h.next != nil {
		return h.next.Handle(form)
	}
	return nil
}

// Обработчик проверки на пустые поля
type NotEmptyHandler struct {
	BaseHandle
}

// Конструктор обработчика NotEmptyHandler
func NewNotEmptyHandler() *NotEmptyHandler {
	return &NotEmptyHandler{}
}

// Реализация проверки: если поля пустые, возвращает ошибку
func (h *NotEmptyHandler) Handle(form *Registration) error {
	if form.email == "" || form.password == "" {
		return &ValidationError{Message: "Error: Email or password cannot be empty."}
	}
	return h.BaseHandle.Handle(form)
}

// Обработчик проверки длины пароля
type PasswordHandler struct {
	BaseHandle
}

// Конструктор обработчика PasswordHandler
func NewPasswordHandler() *PasswordHandler {
	return &PasswordHandler{}
}

// Реализация проверки длины пароля: если меньше 8 символов, возвращает ошибку
func (h *PasswordHandler) Handle(form *Registration) error {
	if len(form.password) < 8 {
		return &ValidationError{Message: "Error: Password must be at least 8 characters long."}
	}
	return h.BaseHandle.Handle(form)
}

// Обработчик проверки формата email
type EmailFormatHandler struct {
	BaseHandle
}

// Конструктор обработчика EmailFormatHandler
func NewEmailFormatHandler() *EmailFormatHandler {
	return &EmailFormatHandler{}
}

// Реализация проверки формата email: если нет "@" в адресе, возвращает ошибку
func (h *EmailFormatHandler) Handle(form *Registration) error {
	if !strings.Contains(form.email, "@") {
		return &ValidationError{Message: "Error: Invalid email format."}
	}
	return h.BaseHandle.Handle(form)
}

func main() {
	form := &Registration{email: "test@gmail.com", password: "12345678"}

	notEmptyHandler := NewNotEmptyHandler()
	passwordHandler := NewPasswordHandler()
	emailFormatHandler := NewEmailFormatHandler()

	// Построение цепочки обработчиков
	notEmptyHandler.SetNext(passwordHandler).SetNext(emailFormatHandler)

	if err := notEmptyHandler.Handle(form); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Registration completed successfully")
	}
}

// Паттерн "Цепочка вызовов" позволяет последовательно вызывать методы объектов,
// передавая управление от одного объекта к другому, не привязываясь к конкретной
// реализации каждого из объектов в Go.
// Паттерн "Цепочка вызовов" позволяет выстраивать последовательные вызовы
// методов, где каждый объект может обработать запрос или передать его
// следующему в цепочке объекту.
// Это позволяет динамически управлять обработкой запроса, добавляя или
// удаляя звенья в цепочке без изменения основной логики.

// Плюсы применения паттерна "Цепочка вызовов":
// 1. Уменьшает зависимость между клиентом и обработчиками.
//    Паттерн позволяет клиенту не знать, какой конкретный обработчик будет выполнять запрос,
//    так как цепочка обработчиков скрывает детали реализации каждого обработчика.
//
// 2. Реализует принцип единственной обязанности.
//    Каждый обработчик выполняет одну конкретную задачу (например, проверяет email или пароль),
//    что делает код более чистым и разделяет ответственность.
//
// 3. Реализует принцип открытости/закрытости.
//    Добавление новых обработчиков в цепочку не требует изменения существующих,
//    что делает код более расширяемым и менее подверженным изменениям.
//
// 4. Запрос может остаться никем не обработанным.
//    Если ни один из обработчиков не может обработать запрос, он просто пройдет через всю цепочку,
//    и результатом будет отсутствие обработки, что позволяет гибко управлять запросами.

// Минусы применения паттерна "Цепочка вызовов":
// 1. Может быть сложным для отладки.
//    Если цепочка обработчиков длинная и сложная, может быть трудно проследить,
//    на каком этапе и по какой причине запрос не был обработан.
//
// 2. Может привести к излишнему числу обработчиков.
//    Если цепочка обработки запросов становится слишком длинной,
//    это может привести к тому, что код станет трудным для понимания и сопровождения.
//
// 3. Производительность может пострадать при очень длинных цепочках.
//    Если цепочка содержит слишком много звеньев, это может замедлить обработку запроса,
//    так как каждый обработчик будет последовательно проверять и передавать запрос дальше.
