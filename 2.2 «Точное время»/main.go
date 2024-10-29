package main

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

// Программа запрашивает текущее время у NTP-сервера "pool.ntp.org".
// Если запрос успешен, программа выводит полученное время.
// В случае ошибки, программа выводит сообщение об ошибке в STDERR
// и завершает выполнение с ненулевым кодом выхода.
func main() {
	currentTime, err := ntp.Time("pool.ntp.org")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error retrieving accurate time:", err)
		os.Exit(1)
	}

	fmt.Println("Accurate time:", currentTime)
}
