package main

import "fmt"

// Задача 1.1
// Объявить переменные всех основных типов, 3 способами: var с типом, var без типа, :=
// Вывести значение через %T

func main() {
	// var с типом
	var a1 int = 10
	var a2 float64 = 10.5
	var a3 string = "Hello"
	var a4 bool = true
	var a5 byte = 255
	var a6 rune = 'a'

	// var без типа
	var b1 = 10
	var b2 = 10.5
	var b3 = "Hello"
	var b4 = true
	var b5 = byte(255)
	var b6 = 'a'

	// :=
	c1 := 10
	c2 := 10.5
	c3 := "Hello"
	c4 := true
	c5 := byte(255)
	c6 := 'a'

	fmt.Printf("%T\n %v\n", a1, a1)
	fmt.Printf("%T\n %v\n", a2, a2)
	fmt.Printf("%T\n %v\n", a3, a3)
	fmt.Printf("%T\n %v\n", a4, a4)
	fmt.Printf("%T\n %v\n", a5, a5)
	fmt.Printf("%T\n %v\n", a6, a6)

	fmt.Printf("%T\n %v\n", b1, b1)
	fmt.Printf("%T\n %v\n", b2, b2)
	fmt.Printf("%T\n %v\n", b3, b3)
	fmt.Printf("%T\n %v\n", b4, b4)
	fmt.Printf("%T\n %v\n", b5, b5)
	fmt.Printf("%T\n %v\n", b6, b6)

	fmt.Printf("%T\n %v\n", c1, c1)
	fmt.Printf("%T\n %v\n", c2, c2)
	fmt.Printf("%T\n %v\n", c3, c3)
	fmt.Printf("%T\n %v\n", c4, c4)
	fmt.Printf("%T\n %v\n", c5, c5)
	fmt.Printf("%T\n %v\n", c6, c6)
}