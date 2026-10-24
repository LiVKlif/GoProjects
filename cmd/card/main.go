package main

import "fmt"

func main() {
	const course = "Бэкенд-разработка на языке Go"
	var name string = "Гугл"
	var group string = "ТЭП-03-26"
	completed := 1
	var ready bool = true

	fmt.Println("Участник:", name)
	fmt.Printf("Группа: %s\n", group)
	fmt.Printf("Курс: %s\n", course)
	fmt.Printf("Завершено работ: %d\n", completed)
	fmt.Printf("Интерес к бэкенду: %t\n", ready)
}
