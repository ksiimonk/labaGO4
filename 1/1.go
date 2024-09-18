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
	}

	// Добавляем нового человека
	people["Мария"] = 28

	// Выводим все записи на экран
	for name, age := range people {
		fmt.Printf("Имя: %s, Возраст: %d\n", name, age)
	}
}
