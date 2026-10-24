package main

import "fmt"

func main() {
	var n int
	fmt.Print("Введите N: ")
	fmt.Scan(&n)
	if n < 1 {
		fmt.Println("N должно быть >= 1")
		return
	}
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	avg := float64(sum) / float64(n)
	fmt.Printf("Сумма чисел от 1 до %d = %d\n", n, sum)
	fmt.Printf("Среднее значение: %.2f\n", avg)
}
