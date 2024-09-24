package main

import "fmt"

func checker(num int) string {
	if num == 0 {
		return "Zero"
	}
	if num > 0 {
		return "Positive"
	}
	if num < 0 {
		return "Negative"
	}
	return "Ошибка"

}

func stringLength(s string) int {
	return len(s)
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func calculateAverage(a int, b int) float64 {
	return float64(a+b) / 2.0
}

func main() {
	var num int

	fmt.Print("Введите число: ")
	fmt.Scanln(&num)

	fmt.Println("Задание 1")
	if num%2 == 0 {
		fmt.Println("Число", num, "является четным.")
	} else {
		fmt.Println("Число", num, "является нечетным.")
	}

	fmt.Println("Задание 2")
	fmt.Println(checker(num))

	fmt.Println("Задание 3")
	for i := 1; i < 11; i++ {
		fmt.Println(i)
	}

	fmt.Println("Задание 4")
	str := "Hi there!"
	length := stringLength(str)
	fmt.Println("Длина строки:", length, "символов")

	fmt.Println("Задание 5")
	rect := Rectangle{Width: 5.12, Height: 2.03}
	fmt.Println("Площадь прямоугольника:", rect.Area())

	fmt.Println("Задание 6")
	num1 := 5
	num2 := 10
	average := calculateAverage(num1, num2)
	fmt.Println("Среднее значение:", average)
}
