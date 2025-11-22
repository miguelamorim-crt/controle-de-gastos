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
	fmt.Println("3 - lista de compras")
	fmt.Println("4 - ver salario")
	fmt.Println("5 - definir ou redefinir salario mensal(crucial)")
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

	//tratamento de entrada simples
	if quantidade < 0 {
		fmt.Println("erro: quantidade menor que 0")
		return
	}

	if salario == 0 {
		fmt.Println("erro: nao definiu o salario")
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

		//tratamento de entrada basico
		if ganha_em_mes < 0 {
			fmt.Println("valor menor que 0!")
			return
		}

		salario = ganha_em_mes
		fmt.Println("definido!")
		definiu = true
	} else {
		var antigo float64
		fmt.Print("redefina seu salario mensal: ")
		antigo = salario
		fmt.Scanln(&ganha_em_mes)

		//tratamento de entrada basico
		if ganha_em_mes < 0 {
			fmt.Println("valor menor que 0!")
			return
		}

		salario = ganha_em_mes
		fmt.Println("novo salario: ", salario)
		fmt.Println("antigo salario: ", antigo)
	}
}

func compras() {
	var nome_produto string
	var preco_produto float64
	var quantidade_compras int
	var preco_final_compras float64

	fmt.Println("coloque itens na sua lista")

	for {
		fmt.Print("qual o nome do produto: ")
		fmt.Scanln(&nome_produto)
		fmt.Print("qual o valor do produto(0.00 para encerrar a lista de compras): ")
		fmt.Scanln(&preco_produto)

		fmt.Println("==============================================================")

		if preco_produto == 0.00 {
			break
		}

		quantidade_compras += 1

		preco_final_compras += preco_produto
	}
	fmt.Println("===  ESTASTISTICAS DA COMPRA ===")
	fmt.Println("itens: ", quantidade_compras)
	fmt.Println("preço total dos produtos: ", preco_final_compras)
}

func separador() {
	fmt.Println("================================")
}

func main() {
	var opcao int
	for {
		menu()
		fmt.Scanln(&opcao)

		separador()
		switch opcao {
		case 0:
			fmt.Println("saindo...")
			return
		case 1:
			adicionarGastos()
		case 2:
			verGastos()
		case 3:
			compras()
		case 4:
			fmt.Println("salario: ", salario)
		case 5:
			redefinir()
		default:
			fmt.Println("opcao invalida")
		}
		separador()

	}
}

