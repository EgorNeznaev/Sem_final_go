package main

import "fmt"

func main() {
	var balance float64
	fmt.Print("Введите начальный баланс: ")
	fmt.Scan(&balance)
	var choice int
	fmt.Println("Выберите действие:")
	fmt.Println("1 - Посмотреть баланс")
	fmt.Println("2 - Пополнить баланс")
	fmt.Println("3 - Снять средства")
	fmt.Print("Ваш выбор: ")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		fmt.Printf("Текущий баланс: %.2f\n", balance)
	case 2:
		var amount float64
		fmt.Print("Введите сумму пополнения: ")
		fmt.Scan(&amount)
		if amount <= 0 {
			fmt.Println("Сумма должна быть положительной")
			return
		}
		balance += amount
		fmt.Printf("Баланс пополнен. Новый баланс: %.2f\n", balance)
	case 3:
		var amount float64
		fmt.Print("Введите сумму для снятия: ")
		fmt.Scan(&amount)
		if amount <= 0 {
			fmt.Println("Сумма должна быть положительной")
			return
		}
		if amount > balance {
			fmt.Println("Недостаточно средств на балансе")
		} else {
			balance -= amount
			fmt.Printf("Средства сняты. Новый баланс: %.2f\n", balance)
		}
	default:
		fmt.Println("Неверный выбор")
	}
}
