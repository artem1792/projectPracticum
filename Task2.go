// По даному трехзначному числу определите, все ли его цифры различны. 
// Формат входных данных: на вход подается одно натуральное трехзначное число. 
// Формат выходных данных: выведите "YES", если все цифры числа различны, в противном случае - "NO".
package main

import "fmt"

func main() {
	fmt.Println(Task2(172))
}

func Task2(number int) string {
	s := number / 100
	d := (number / 10) % 10
	e := number % 10

	if s != d && d != e && s != e {
		return "YES"
	} else {
		return "NO"
	}
}