package calculator

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/main.go/roman"
)

const (
	Sum            = "+"
	Multiplication = "*"
	Division       = "/"
	Subtraction    = "-"
)

type Calculator struct {
	Input         string
	MathOperation string
	FirstValue    int
	SecondValue   int
}

func (c *Calculator) ParseInput() error {
	parts := strings.Fields(c.Input)
	if len(parts) != 3 {
		return errors.New("неверный формат ввода")
	}

	firstValueStr := parts[0]
	c.MathOperation = parts[1]
	secondValueStr := parts[2]
	var err error
	var isFirstRoman bool
	var isSecondRoman bool

	c.FirstValue, isFirstRoman = roman.RomanToArabic(firstValueStr)
	c.SecondValue, isSecondRoman = roman.RomanToArabic(secondValueStr)

	if isFirstRoman == true && isSecondRoman == true {
		return nil
	}

	if isFirstRoman != isSecondRoman {
		fmt.Println("неправильный формат ввода")
		return nil
	}

	c.FirstValue, err = strconv.Atoi(firstValueStr) // НЕ ДОПИСАНО C
	if err != nil {
		fmt.Println("Ошибка при преобразовании первого значения", err)
		return nil

	}

	c.SecondValue, err = strconv.Atoi(secondValueStr)
	if err != nil {
		fmt.Println("Ошибка при преобразовании второго значения", err)
		return nil
	}

	return nil

}
func NewCalculator(input string) *Calculator {
	calculator := &Calculator{Input: input}
	err := calculator.ParseInput()
	if err != nil {
		return nil

	}

	return calculator
}

func (c *Calculator) Calculate() int {
	switch c.MathOperation {
	case Sum:
		return c.FirstValue + c.SecondValue
	case Subtraction:
		return c.FirstValue - c.SecondValue
	case Multiplication:
		return c.FirstValue * c.SecondValue
	case Division:
		if c.SecondValue == 0 {
			fmt.Println("Ошибка деление на ноль")
			return 0
		}
		return c.FirstValue / c.SecondValue
	default:
		fmt.Println("Неверная математическая операция")
		return -1
	}
}
