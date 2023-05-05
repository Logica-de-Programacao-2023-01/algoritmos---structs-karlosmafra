package main

import "fmt"

type Endereco struct {
	rua    string
	num    int
	cidade string
	estado string
}

type Pessoa struct {
	nome  string
	idade int
	end   Endereco
}

func main() {
	p := Pessoa{
		nome:  "",
		idade: 20,
		end: Endereco{
			rua:    "Rua X",
			num:    50,
			cidade: "Brasília",
			estado: "DF",
		}}

	imprimirEnd(p)
}

func imprimirEnd(p Pessoa) {
	fmt.Println(p.end.rua)
	fmt.Println(p.end.num)
	fmt.Println(p.end.cidade)
	fmt.Println(p.end.estado)
}
