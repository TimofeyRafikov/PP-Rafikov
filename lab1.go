package main

import (
	"fmt"
	"time"
)

// Функция для вычисления суммы и разности двух чисел с плавающей запятой
func sumAndDifference(a, b float64) (float64, float64) {
	sum := a + b
	diff := a - b
	return sum, diff
}

// Функция для вычисления среднего значения трех чисел
func average(a, b, c float64) float64 {
	return (a + b + c) / 3
}

func main() {
	// Задача 1: Вывод текущей даты и времени
	fmt.Println("Задача 1: Текущая дата и время")
	currentTime := time.Now()
	fmt.Println("Текущее время:", currentTime.Format("2006-01-02 15:04:05"))
	fmt.Println()

	// Задача 2: Создание переменных различных типов и вывод их на экран
	fmt.Println("Задача 2: Переменные различных типов")
	var i int = 10
	var f float64 = 3.14
	var s string = "Привет, мир!"
	var b bool = true
	fmt.Printf("int: %d, float64: %.2f, string: %s, bool: %t\n", i, f, s, b)
	fmt.Println()

	// Задача 3: Использование краткой формы объявления переменных
	fmt.Println("Задача 3: Краткая форма объявления переменных")
	a := 5
	c := 2.5
	d := "Go"
	e := false
	fmt.Printf("a: %d, c: %.2f, d: %s, e: %t\n", a, c, d, e)
	fmt.Println()

	// Задача 4: Арифметические операции с двумя числами
	fmt.Println("Задача 4: Арифметические операции")
	x, y := 7, 3
	fmt.Printf("%d + %d = %d\n", x, y, x+y)
	fmt.Printf("%d - %d = %d\n", x, y, x-y)
	fmt.Printf("%d * %d = %d\n", x, y, x*y)
	fmt.Printf("%d / %d = %d\n", x, y, x/y)
	fmt.Println()

	// Задача 5: Функция для вычисления суммы и разности двух чисел с плавающей запятой
	fmt.Println("Задача 5: Сумма и разность двух чисел с плавающей запятой")
	num1, num2 := 4.5, 1.2
	sum, diff := sumAndDifference(num1, num2)
	fmt.Printf("Сумма: %.2f, Разность: %.2f\n", sum, diff)
	fmt.Println()

	// Задача 6: Вычисление среднего значения трех чисел
	fmt.Println("Задача 6: Среднее значение трех чисел")
	avg := average(3.0, 4.0, 5.0)
	fmt.Printf("Среднее значение: %.2f\n", avg)
}
