package main

import "fmt"

// Código relacionado con la función Principal
func Principal() string {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	return "hola"
}
