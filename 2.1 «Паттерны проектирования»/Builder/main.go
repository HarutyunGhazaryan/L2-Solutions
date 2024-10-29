package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Интерфейс Builder для гибкости
type Builder interface {
	Add(key string, value interface{}) Builder
	AddNested(key string, nestedBuilder Builder) Builder
	Build() (string, error)
}

// JSONBuilder реализует интерфейс Builder
type JSONBuilder struct {
	data map[string]interface{}
}

// Фабричный метод для создания нового строителя JSON
func NewBuilder() *JSONBuilder {
	return &JSONBuilder{
		data: make(map[string]interface{}),
	}
}

// Добавляем пару ключ-значение в структуру
func (b *JSONBuilder) Add(key string, value interface{}) Builder {
	if _, exist := b.data[key]; exist {
		log.Printf("Warning: Key %s already exists and will be overwritten.\n", key)
	}
	b.data[key] = value
	return b
}

// Добавляем вложенную структуру JSON, принимая интерфейс Builder
func (b *JSONBuilder) AddNested(key string, nestedBuilder Builder) Builder {
	nestedJSON, err := nestedBuilder.Build()
	if err != nil {
		log.Printf("Error building nested JSON for key %s: %v", key, err)
	}
	var nestedData map[string]interface{}
	err = json.Unmarshal([]byte(nestedJSON), &nestedData)
	if err != nil {
		log.Printf("Error unmarshaling nested JSON for key %s: %v", key, err)
	}
	b.data[key] = nestedData
	return b
}

// Собираем JSON-строку из данных
func (b *JSONBuilder) Build() (string, error) {
	jsonData, err := json.Marshal(b.data)
	if err != nil {
		return "", fmt.Errorf("error marshaling JSON: %w", err)
	}
	return string(jsonData), nil
}

// Директор управляет процессом сборки
type Director struct {
	builder Builder
}

// Конструктор для директора
func NewDirector(b Builder) *Director {
	return &Director{builder: b}
}

// Строим простой JSON
func (d *Director) BuildSimpleJSON() (string, error) {
	d.builder.Add("name", "Name1").Add("age", 20)
	return d.builder.Build()
}

// Строим сложный JSON
func (d *Director) BuildComplexJSON() (string, error) {
	d.builder.Add("name", "Name1").Add("age", 20).Add("preferences", map[string]interface{}{"notification": true})
	return d.builder.Build()
}

func main() {
	// Создаем строителя и директора
	builder := NewBuilder()
	director := NewDirector(builder)

	// Строим простой JSON
	simpleJSON, err := director.BuildSimpleJSON()
	if err != nil {
		log.Printf("Error building simple JSON: %v", err)
	} else {
		fmt.Println("Simple JSON:", simpleJSON)
	}

	// Строим сложный JSON
	complexJSON, err := director.BuildComplexJSON()
	if err != nil {
		log.Printf("Error building complex JSON: %v", err)
	} else {
		fmt.Println("Complex JSON:", complexJSON)
	}

	// Пример добавления вложенной структуры с телефоном
	phoneBuilder := NewBuilder().Add("phone", "123456")
	builder.AddNested("contact", phoneBuilder)

	// Строим JSON с контактными данными
	nestedJSON, err := builder.Build()
	if err != nil {
		log.Printf("Error building JSON with contact data: %v", err)
	} else {
		fmt.Println("Nested JSON:", nestedJSON)
	}
}

// Применение паттерна Строитель (Builder) в данном коде заключается в гибком
// и пошаговом создании сложных объектов (в данном случае JSON). Он позволяет
// отделить процесс создания объекта от его структуры и позволяет легко
// управлять добавлением новых полей или вложенных объектов.
// В данном примере JSONBuilder используется для динамического создания JSON-объектов,
// а Director управляет процессом сборки для упрощения создания типичных конфигураций JSON.

// Плюсы паттерна:
// 1. Гибкость: позволяет пошагово добавлять поля в структуру и легко создавать
//    как простые, так и сложные объекты.
// 2. Читабельность: код становится более читабельным благодаря цепочке вызовов методов
//    (chaining), что улучшает понимание процесса сборки объекта.
// 3. Возможность переиспользования: можно легко создать различные конфигурации объекта,
//    используя одного и того же строителя (например, BuildSimpleJSON и BuildComplexJSON).
// 4. Контроль за процессом: использование директора (Director) помогает управлять процессом
//    сборки и предлагать готовые шаблоны для создания объектов.

// Минусы паттерна:
// 1. Усложнение кода: добавление интерфейсов и классов (Builder, Director) может увеличить
//    сложность структуры кода, особенно если требуется создание простых объектов.
// 2. Многословность: для небольших или простых объектов, процесс сборки может оказаться излишне
//    сложным, что может увеличить количество кода.
// 3. Производительность: из-за цепочек вызовов и возможного логирования могут возникнуть
//    дополнительные накладные расходы на производительность.
