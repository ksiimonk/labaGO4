package main

import (
	"fmt"
)

func main() {
	// Создаем карту с именами и возрастами
	people := map[string]int{
		"Иван":   30,
		"Анна":   25,
		"Сергей": 40,
		"Мария":  28,
	}

	// Имя для удаления
	nameToDelete := "Анна"

	// Удаляем запись
	delete(people, nameToDelete)

	// Выводим оставшиеся записи
	for name, age := range people {
		fmt.Printf("Имя: %s, Возраст: %d\n", name, age)
	}
}
