package main

import (
	"fmt"
)

func main() {
	n := 0
	fmt.Print("Введите целое число: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Вы ввели число: %d\n", n)
}
