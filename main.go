package main

import (
	"fmt"
	"log"
)

func main() {
	var n interface{}
	fmt.Print("Введите целое число: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Вы ввели число: %d\n", n)
}
