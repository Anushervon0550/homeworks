package main

import (
	"fmt"
	"math"
)

// 1
type Book struct {
	Title  string
	Author string
}

func (b Book) GetDetails() string {
	return fmt.Sprintf("Title: %s, Author: %s", b.Title, b.Author)
}

type Library struct {
	Books []Book
}

func (l *Library) AddBook(book Book) {
	l.Books = append(l.Books, book)
}

// 2
type Student struct {
	Grades []int
}

func (s *Student) AddGrade(grade int) {
	s.Grades = append(s.Grades, grade)
}
func (s Student) AverageGrade() float64 {
	if len(s.Grades) == 0 {
		return 0
	}
	sum := 0
	for _, grade := range s.Grades {
		sum += grade
	}
	return float64(sum) / float64(len(s.Grades))
}

// 3
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (c Circle) Circumference() float64 {
	return 2 * math.Pi * c.Radius
}

func main() {
	//1
	book1 := Book{Title: "Go Programming", Author: "John Doe"}
	book2 := Book{Title: "Advanced Go", Author: "Jane Smith"}
	library := Library{}
	library.AddBook(book1)
	library.AddBook(book2)
	for _, book := range library.Books {
		fmt.Println(book.GetDetails())
	}
	//2
	// Создание экземпляра студента
	student := Student{}

	// Добавление оценок
	student.AddGrade(85)
	student.AddGrade(90)
	student.AddGrade(78)

	// Вывод среднего значения оценок
	fmt.Printf("Average Grade: %.2f\n", student.AverageGrade())
	//3
	circle := Circle{Radius: 5}

	// Вывод площади и периметра круга
	fmt.Printf("Area: %.2f\n", circle.Area())
	fmt.Printf("Circumference: %.2f\n", circle.Circumference())
}
