package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите числа через пробел: ")
	input, _ := reader.ReadString('\n')
	// Разделяем строку на части
	numbersStr := strings.Fields(input)
	numbers := make([]int, len(numbersStr))

	for i, numStr := range numbersStr {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("Ошибка преобразования числа:", err)
			return
		}
		numbers[i] = num
	}

	// Выводим массив в обратном порядке
	for i := len(numbers) - 1; i >= 0; i-- {
		fmt.Printf("%d ", numbers[i])
	}
	fmt.Println()
}
