package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите строку: ")
	input, _ := reader.ReadString('\n')
	// Удаляем символ новой строки
	input = strings.TrimSpace(input)
	// Выводим строку в верхнем регистре
	fmt.Println(strings.ToUpper(input))
}
