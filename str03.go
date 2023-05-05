package main

import "fmt"

type Triangulo struct {
	base   float64
	altura float64
}

func main() {
	var t Triangulo

	fmt.Print("Base: ")
	fmt.Scan(&t.base)
	fmt.Print("Altura: ")
	fmt.Scan(&t.altura)

	fmt.Println("A área do triângulo é", areaT(t))
}

func areaT(t Triangulo) float64 {
	area := t.base * t.altura / 2
	return area
}
