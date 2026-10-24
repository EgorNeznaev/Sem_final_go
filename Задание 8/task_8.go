package main

import "fmt"

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func main() {
	var n int
	fmt.Print("Введите число для проверки: ")
	fmt.Scan(&n)
	if isPrime(n) {
		fmt.Printf("%d — простое число\n", n)
	} else {
		fmt.Printf("%d — не является простым числом\n", n)
	}
}
