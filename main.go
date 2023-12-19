package main

import (
	"fmt"
)

func main() {
	var n interface{}
	fmt.Println("Введите целое число: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Вы ввели число: %d\n", n)
}
