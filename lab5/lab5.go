package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	name string
	age  int
}

func (p Person) DisplayInfo() {
	fmt.Printf("Имя: %s, Возраст: %d лет\n", p.name, p.age)
}

func (p *Person) Birthday() {
	p.age++
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func printAreas(shapes []Shape) {
	for _, shape := range shapes {
		fmt.Printf("Площадь: %.2f\n", shape.Area())
	}
}

type Stringer interface {
	String() string
}

type Book struct {
	Title  string
	Author string
	Year   int
}

func (b Book) String() string {
	return fmt.Sprintf("Книга: %s, Автор: %s, Год: %d", b.Title, b.Author, b.Year)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	person1 := Person{name: "Pasha", age: 20}
	person2 := Person{name: "Dima", age: 100}

	person1.DisplayInfo()
	person2.DisplayInfo()

	fmt.Print("Хотите увеличить возраст Pasha на 1 год? (да/нет): ")
	scanner.Scan()
	input := scanner.Text()
	if strings.ToLower(input) == "да" {
		person1.Birthday()
	}

	fmt.Print("Хотите увеличить возраст Dima на 1 год? (да/нет): ")
	scanner.Scan()
	input = scanner.Text()
	if strings.ToLower(input) == "да" {
		person2.Birthday()
	}
	person1.DisplayInfo()
	person2.DisplayInfo()

	//2

	scanner1 := bufio.NewScanner(os.Stdin)

	fmt.Print("Введите радиус круга: ")
	scanner1.Scan()
	input1 := scanner1.Text()

	radius, err := strconv.ParseFloat(input1, 64)
	if err != nil {
		fmt.Println("Ошибка: неверный формат радиуса")
		return
	}

	circle := Circle{radius: radius}

	area := circle.Area()
	fmt.Printf("Площадь круга с радиусом %.2f: %.2f\n", circle.radius, area)

	//3

	rectangle := Rectangle{width: 5.0, height: 10.0}

	shapes := []Shape{rectangle, circle}

	printAreas(shapes)

	//4

	book := Book{Title: "1984", Author: "Джордж Оруелл", Year: 1949}

	fmt.Println(book.String())

}
