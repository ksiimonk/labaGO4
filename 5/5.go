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
	sum := 0

	for _, numStr := range numbersStr {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("Ошибка преобразования числа:", err)
			return
		}
		sum += num
	}

	// Выводим сумму
	fmt.Printf("Сумма чисел: %d\n", sum)
}
