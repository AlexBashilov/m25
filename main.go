package main

import (
	"fmt"
	"log"
)

func main() {
	n := 0
	fmt.Print("Введите данные: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Вы ввели следующие данные: %d\n", n)
}
