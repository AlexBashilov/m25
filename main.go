package main

import (
	"fmt"
)

func main() {
	var n interface{}
	fmt.Print("Введите данные: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Вы ввели следующие данные: %d\n", n)
}
