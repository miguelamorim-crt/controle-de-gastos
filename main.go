package main

import "fmt"

var definiu bool
var salario float64
var gastos []string
var valores []float64

//menu
func menu() {
	fmt.Println("===  CONTROLE DE GASTOS ===")
	fmt.Println("1 - adicionar gasto")
	fmt.Println("2 - ver gastos")
	fmt.Println("3 - definir ou redefinir salario mensal(crucial)")
	fmt.Println("0 - sair")
	fmt.Println("===========================")
	fmt.Print("opcao escolhida: ")
}

//adiciona o gasto
func adicionarGastos() {
	var nome string
	var custo float64

	fmt.Print("quantos gastos quer adicionar: ")
	var quantidade int
	fmt.Scanln(&quantidade)

	if salario == 0 {
		fmt.Println("parece que voce nao definiu seu salario!")
		return
	}
	for i := 1; i <= quantidade; i++ {
		fmt.Println("gasto numero ", i)

		fmt.Print("qual o nome: ")
		fmt.Scanln(&nome)

		fmt.Print("quantos gasta com isso mensalmente: ")
		fmt.Scanln(&custo)

		salario -= custo
		fmt.Println("salario depois da alteraçao: ", salario)

		gastos = append(gastos, nome)
		valores = append(valores, custo)
	}
	fmt.Println("salario final: ", salario)
}

func verGastos() {
	fmt.Println("salario: ", salario)
	for i := 0; i < len(gastos); i++ {
		fmt.Printf("%d. nome: %s | gasto: R$ %.2f\n", i+1, gastos[i], valores[i])
	}
}

//define ou redefine o salario do usuario
func redefinir() {
	var ganha_em_mes float64
	if definiu == false {
		fmt.Print("defina seu salario mensal(quantos ganha por mes): ")
		fmt.Scanln(&ganha_em_mes)
		salario = ganha_em_mes
		fmt.Println("definido!")
		definiu = true
	} else {
		var antigo float64
		fmt.Print("redefina seu salario mensal: ")
		antigo = salario
		fmt.Scanln(&ganha_em_mes)
		salario = ganha_em_mes

		fmt.Println("novo salario: ", salario)
		fmt.Println("antigo salario: ", antigo)
	}
}

func main() {
	var opcao int
	for {
		menu()
		fmt.Scanln(&opcao)

		fmt.Println("===========================")
		switch opcao {
		case 0:
			fmt.Println("saindo...")
			return
		case 1:
			adicionarGastos()
		case 2:
			verGastos()
		case 3:
			redefinir()
		default:
			fmt.Println("opcao invalida")
		}
		fmt.Println("===========================")

	}
}
