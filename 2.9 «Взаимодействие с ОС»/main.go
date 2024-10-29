package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// getCurrentDir возвращает текущую рабочую директорию.
// Если возникает ошибка при получении директории, она будет выведена в консоль.
func getCurrentDir() string {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
	}
	return currentDir
}

// printPrompt выводит приглашение с текущей директорией для ввода команды.
func printPrompt() {
	currentDir := getCurrentDir()
	fmt.Print(currentDir + "> ")
}

// executeCommand выполняет введённую команду, обрабатывая основные команды оболочки.
// В зависимости от команды, она может:
// - Изменять директорию (cd)
// - Печать текущей директории (pwd)
// - Выводить текст (echo)
// - Убивать процессы (kill)
// - Показывать запущенные процессы (ps)
// - Выполнять другие команды, переданные пользователем
func executeCommand(command string) {
	args := strings.Fields(command)
	if len(args) == 0 {
		return
	}

	switch args[0] {
	case "cd":
		if len(args) > 1 {
			err := os.Chdir(args[1])
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			err := os.Chdir(os.Getenv("HOME"))
			if err != nil {
				fmt.Println("Error:", err)
			}
		}
	case "pwd":
		currentDir := getCurrentDir()
		fmt.Println(currentDir)
	case "echo":
		if len(args) > 1 {
			fmt.Println(strings.Join(args[1:], " "))
		}
	case "kill":
		if len(args) < 2 {
			fmt.Println("Usage: kill <pid>")
			break
		}

		pid, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Invalid PID:", args[1])
			break
		}

		process, err := os.FindProcess(pid)
		if err != nil {
			fmt.Println("Error finding process:", err)
			break
		}

		err = process.Kill()
		if err != nil {
			fmt.Println("Error killing process:", err)
		} else {
			fmt.Println("Process", pid, "terminated")
		}
	case "ps":
		cmd := exec.Command("powershell", "Get-Process")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println("Error executing ps:", err)
		}
	case "exit", "\\quit":
		fmt.Println("Exiting shell...")
		os.Exit(0)
	default:
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}

// executePipeline выполняет последовательные команды, переданные через конвейер (пайп).
// Каждая команда получает вывод предыдущей команды в качестве ввода.
func executePipeline(commands []string) {
	var lastCmd *exec.Cmd
	for i, cmdStr := range commands {
		args := strings.Fields(cmdStr)
		if len(args) == 0 {
			continue
		}

		if i == 0 {
			lastCmd = exec.Command(args[0], args[1:]...)
		} else {
			nextCmd := exec.Command(args[0], args[1:]...)
			if lastCmd != nil {
				// Создание пайпа
				pipe, err := lastCmd.StdoutPipe()
				if err != nil {
					fmt.Println("Error creating pipe:", err)
					return
				}
				nextCmd.Stdin = pipe
			}
			lastCmd = nextCmd
		}

		if i > 0 {
			lastCmd.Stdout = os.Stdout
		}
	}

	if lastCmd != nil {
		if err := lastCmd.Start(); err != nil {
			fmt.Println("Error starting command:", err)
			return
		}
		lastCmd.Wait()
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	printPrompt()

	for {
		if scanner.Scan() {
			input := scanner.Text()
			commands := strings.Split(input, "|")

			if len(commands) > 1 {
				executePipeline(commands)
			} else {
				executeCommand(input)
			}

			printPrompt()
		} else {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}
