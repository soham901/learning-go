package main

import (
	"fmt"
	"log"

	"example.com/calc"
)

func main() {
	log.SetPrefix("MY LOGS: ")
	log.SetFlags(0)

	res := calc.Add(2, 5)
	fmt.Println(res)

	res = calc.Mul(2, 5)
	fmt.Println(res)

	message, err := calc.Hello("Soham")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	msg, err := calc.Hello("")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)
}
