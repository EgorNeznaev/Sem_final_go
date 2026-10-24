package main

import (
	"fmt"
)

func main() {
	var a, b float64
	var op string
	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)
	fmt.Print("Введите операцию (+, -, *, /): ")
	fmt.Scan(&op)
	switch op {
	case "+":
		fmt.Printf("%.2f + %.2f = %.2f\n", a, b, a+b)
	case "-":
		fmt.Printf("%.2f - %.2f = %.2f\n", a, b, a-b)
	case "*":
		fmt.Printf("%.2f * %.2f = %.2f\n", a, b, a*b)
	case "/":
		if b == 0 {
			fmt.Println("Деление на ноль")
		} else {
			fmt.Printf("%.2f / %.2f = %.2f\n", a, b, a/b)
		}
	default:
		fmt.Println("Неизвестная операция")
	}
}
