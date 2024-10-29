package main

import "fmt"

type Command interface {
	Execute()
}

type Cart struct {
	items []string
}

// Метод для добавления товара в корзину
func (c *Cart) AddItem(item string) {
	c.items = append(c.items, item)
	fmt.Printf("Item %s added to the cart\n", item)
}

// Метод для удаления товара из корзины
func (c *Cart) RemoveItem(item string) {
	for index, cartItem := range c.items {
		if cartItem == item {
			c.items = append(c.items[:index], c.items[index+1:]...)
			fmt.Printf("Item %s removed from the cart\n", item)
			return
		}
	}
	fmt.Printf("Item %s not found in the cart\n", item)
}

type AddItemCommand struct {
	cart *Cart
	item string
}

func (a *AddItemCommand) Execute() {
	a.cart.AddItem(a.item)
}

type RemoveItemCommand struct {
	cart *Cart
	item string
}

func (r *RemoveItemCommand) Execute() {
	r.cart.RemoveItem(r.item)
}

type StoreApp struct {
	commandHistory []Command
}

func (s *StoreApp) SetCommand(c Command) {
	s.commandHistory = append(s.commandHistory, c)
	c.Execute()
}

func main() {
	cart := &Cart{}

	addItemCommand := &AddItemCommand{cart: cart, item: "Laptop"}
	addItemCommand2 := &AddItemCommand{cart: cart, item: "Phone"}
	removeItemCommand := &RemoveItemCommand{cart: cart, item: "Phone"}
	removeItemCommand2 := &RemoveItemCommand{cart: cart, item: "Phone"}

	storeApp := &StoreApp{}

	storeApp.SetCommand(addItemCommand)
	storeApp.SetCommand(addItemCommand2)
	storeApp.SetCommand(removeItemCommand)
	storeApp.SetCommand(removeItemCommand2)
}

// Паттерн "Комманда" позволяет инкапсулировать запросы как объекты, чтобы передавать их как параметры, ставя
// абстракцию между вызовом метода и объектом, который его выполняет. В данном примере паттерн используется для
// управления операциями с корзиной (добавление и удаление товаров) без прямого вызова методов корзины.

// Плюсы применения паттерна "Комманда":
// 1. Команды инкапсулируют конкретные действия и позволяют абстрагировать
//    клиентский код от знаний о том, как именно выполняется операция.
// 2. Новые команды можно добавить без необходимости изменения
//    существующего кода.

// Минусы применения паттерна "Комманда":
// 1. Паттерн увеличивает количество типов и может сделать код громоздким, особенно если
//    операций много и они разные по сложности.
// 2. Для каждого действия требуется создавать отдельный объект-команду, что может
//    потребовать дополнительных ресурсов при большом количестве операций.
