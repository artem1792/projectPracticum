// Дано неотрицательное целое число. 
// Найдите и выведите первую цифру числа. 
// Формат входных данных: на вход дается натуральное число, не превосходящее 10000. 
// Формат выходных данных: выведите одно целое число - первую цифру заданного числа.
package main

import "fmt"

func main() {
	fmt.Println(Task3(172))
}

func Task3(number int) string {
	str := fmt.Sprint(number)
	first := string(str[0])

	return first
}