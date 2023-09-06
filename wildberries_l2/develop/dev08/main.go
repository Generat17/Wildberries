package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

/*
=== Взаимодействие с ОС ===

# Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")

		// в зависимости от Line Separator, проверять окончания строк
		if input == "exit\r" {
			return
		}

		args := strings.Split(input, " ")
		switch args[0] {
		case "cd\r":
			if len(args) > 1 {
				err := os.Chdir(args[1])
				if err != nil {
					fmt.Println("cd:", err)
				}
			} else {
				fmt.Println("cd: missing operand")
			}
		case "pwd\r":
			dir, err := os.Getwd()
			if err != nil {
				fmt.Println("pwd:", err)
			} else {
				fmt.Println(dir)
			}
		case "echo\r":
			if len(args) > 1 {
				fmt.Println(strings.Join(args[1:], " "))
			} else {
				fmt.Println()
			}
		case "kill\r":
			if len(args) > 1 {
				cmd := exec.Command("kill", args[1])
				err := cmd.Run()
				if err != nil {
					fmt.Println("kill:", err)
				}
			} else {
				fmt.Println("kill: missing operand")
			}
		case "ps\r":
			cmd := exec.Command("ps")
			output, err := cmd.Output()
			if err != nil {
				fmt.Println("ps:", err)
			} else {
				fmt.Println(string(output))
			}
		default:
			cmd := exec.Command(args[0], args[1:]...)
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout
			err := cmd.Run()
			if err != nil {
				fmt.Println("exec:", err)
			}
		}
	}
}
