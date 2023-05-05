package main

import (
	"fmt"
	"math"
)

type Circulo struct {
	raio float64
}

func main() {
	var c Circulo

	fmt.Print("Informe o raio do círculo: ")
	fmt.Scan(&c.raio)

	fmt.Println("Área do círculo:", areaC(c))
}

func areaC(c Circulo) float64 {
	return math.Pi * c.raio * c.raio
}
