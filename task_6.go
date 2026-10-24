package main

import "fmt"

func main() {
	var n int
	fmt.Print("Введите количество чисел N: ")
	fmt.Scan(&n)
	pos, neg, zero := 0, 0, 0
	for i := 1; i <= n; i++ {
		var num int
		fmt.Printf("Введите число %d: ", i)
		fmt.Scan(&num)
		if num > 0 {
			pos++
		} else if num < 0 {
			neg++
		} else {
			zero++
		}
	}
	fmt.Printf("Положительных: %d\n", pos)
	fmt.Printf("Отрицательных: %d\n", neg)
	fmt.Printf("Нулей: %d\n", zero)
}
