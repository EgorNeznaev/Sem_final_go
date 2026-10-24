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
	sum := 0
	evenCount, oddCount := 0, 0
	var max, min int
	for i := 1; i <= n; i++ {
		var num int
		fmt.Printf("Введите число %d: ", i)
		fmt.Scan(&num)
		sum += num
		if num%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
		if i == 1 {
			max, min = num, num
		} else {
			if num > max {
				max = num
			}
			if num < min {
				min = num
			}
		}
	}
	avg := float64(sum) / float64(n)
	fmt.Printf("Сумма: %d\n", sum)
	fmt.Printf("Среднее: %.2f\n", avg)
	fmt.Printf("Максимум: %d\n", max)
	fmt.Printf("Минимум: %d\n", min)
	fmt.Printf("Чётных: %d\n", evenCount)
	fmt.Printf("Нечётных: %d\n", oddCount)
}
