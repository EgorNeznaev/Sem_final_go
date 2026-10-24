package main

import "fmt"

func main() {
	var n int
	fmt.Print("Введите количество чисел N: ")
	fmt.Scan(&n)
	if n <= 0 {
		fmt.Println("N должно быть положительным")
		return
	}
	var num int
	fmt.Print("Введите число 1: ")
	fmt.Scan(&num)
	max, min := num, num
	for i := 2; i <= n; i++ {
		fmt.Printf("Введите число %d: ", i)
		fmt.Scan(&num)
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	fmt.Printf("Максимальное значение: %d\n", max)
	fmt.Printf("Минимальное значение: %d\n", min)
}
