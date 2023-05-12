package main

import "fmt"

type funcionario struct {
	nom string
	sal float64
	idd int
}

func main() {
	var f1 funcionario

	fmt.Print("Nome do funcionário: ")
	fmt.Scan(&f1.nom)
	fmt.Print("Salário: ")
	fmt.Scan(&f1.sal)
	fmt.Print("Idade: ")
	fmt.Scan(&f1.idd)

	fmt.Println("O novo salário é de", salario(f1))
	fmt.Printf("O tempo de serviço é de %d anos", servico(f1))
}

func salario(f1 funcionario) float64 {
	var m int
	fmt.Print("Digite 1 para aumentar o salário, 2 para diminuir e 3 para manter: ")
	fmt.Scan(&m)
	switch m {
	case 1:
		var p float64
		fmt.Print("Porcentagem de aumento: ")
		fmt.Scan(&p)
		f1.sal += f1.sal / 100 * p
		return f1.sal
	case 2:
		var p float64
		fmt.Print("Porcentagem de diminuição: ")
		fmt.Scan(&p)
		f1.sal -= f1.sal / 100 * p
		return f1.sal
	case 3:
		return f1.sal
	default:
		return 0
	}
}

func servico(f1 funcionario) int {
	return f1.idd - 18
}
