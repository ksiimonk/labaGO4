package main

import (
	"fmt"
)

// Функция для вычисления среднего возраста
func averageAge(people map[string]int) float64 {
	totalAge := 0
	count := 0

	for _, age := range people {
		totalAge += age
		count++
	}

	if count == 0 {
		return 0
	}

	return float64(totalAge) / float64(count)
}

func main() {
	// Создаем карту с именами и возрастами
	people := map[string]int{
		"Иван":   30,
		"Анна":   25,
		"Сергей": 40,
		"Мария":  28,
	}

	// Вычисляем и выводим средний возраст
	fmt.Printf("Средний возраст: %.2f\n", averageAge(people))
}
