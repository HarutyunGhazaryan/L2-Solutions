package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

// main запускает Telnet-клиент, который устанавливает TCP-соединение с указанным
// хостом и портом. Программа считывает ввод пользователя из стандартного ввода (STDIN)
// и отправляет его на сервер, а также выводит полученные данные от сервера в стандартный
// вывод (STDOUT). При необходимости можно указать таймаут для подключения через аргумент
// командной строки. Программа завершает свою работу, если сервер закрывает соединение
// или пользователь нажимает Ctrl+D.
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	timeout := flag.Duration("timeout", 10*time.Second, "Connection timeout")
	flag.Parse()

	if len(flag.Args()) < 2 {
		fmt.Println("Host and port are required")
		os.Exit(1)
	}

	host, port := flag.Arg(0), flag.Arg(1)
	address := net.JoinHostPort(host, port)
	dialer := net.Dialer{Timeout: *timeout}

	conn, err := dialer.Dial("tcp", address)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading from server: %v\n", err)
		}
		cancel()
	}()

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			text := scanner.Text()
			if text == "" {
				continue
			}
			fmt.Fprintln(conn, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading from server: %v\n", err)
		}
		cancel()
	}()

	<-ctx.Done()
}
