package main

import (
	"bufio"

	"fmt"
	"os"
	"strings"

	"github.com/main.go/calculator"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Введите математическую операцию: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Возникла ошибка при чтении ввода: ", err)
			continue
		}
		input = strings.TrimSuffix(input, "\n")
		if input == "exit" {
			fmt.Println("Программа завершена.")
			break
		}
		calc := calculator.NewCalculator(input)
		res := calc.Calculate()
		fmt.Println("Результат: ", res)
	}
}
