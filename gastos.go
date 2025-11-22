package main

import "fmt"

func adicionarGastos() {
    var nome string
    var custo float64

    fmt.Print("quantos gastos quer adicionar: ")
    var quantidade int
    fmt.Scanln(&quantidade)

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
    if len(gastos) == 0 {
        fmt.Println("nao tem nada aqui?")
        return
    }

    fmt.Println("salario: ", salario)
    for i := 0; i < len(gastos); i++ {
        fmt.Printf("%d. nome: %s | gasto: R$ %.2f\n", i+1, gastos[i], valores[i])
    }
}

func removerGasto() {
    if len(gastos) == 0 {
        fmt.Println("erro: nao existe gastos para remover")
        return
    }

    fmt.Println("Lista de gastos:")
    for i := 0; i < len(gastos); i++ {
        fmt.Printf("%d. %s | R$ %.2f\n", i+1, gastos[i], valores[i])
    }

    fmt.Print("Digite o número do gasto que deseja remover: ")
    var indice int
    fmt.Scanln(&indice)

    indice = indice - 1
    if indice < 0 || indice >= len(gastos) {
        fmt.Println("Índice inválido!")
        return
    }

    gastos = append(gastos[:indice], gastos[indice+1:]...)
    valores = append(valores[:indice], valores[indice+1:]...)

    fmt.Println("Gasto removido com sucesso!")
}

func editarGasto() {
    if len(gastos) == 0 {
        fmt.Println("erro: nao existe gastos para remover")
        return
    }

    fmt.Println("Lista de gastos:")
    for i := 0; i < len(gastos); i++ {
        fmt.Printf("%d. %s | R$ %.2f\n", i+1, gastos[i], valores[i])
    }

    fmt.Print("Digite o número do gasto que deseja editar: ")
    var indice int
    fmt.Scanln(&indice)

    indice = indice - 1
    if indice < 0 || indice >= len(gastos) {
        fmt.Println("Índice inválido!")
        return
    }

    var novoNome string
    var novoValor float64

    fmt.Print("novo nome: ")
    fmt.Scanln(&novoNome)
    fmt.Print("novo valor: ")
    fmt.Scanln(&novoValor)

    gastos[indice] = novoNome
    valores[indice] = novoValor

    fmt.Println("Gasto editado com sucesso!")
}

func redefinir() {
    var ganha_em_mes float64
    if definiu == false {
        fmt.Print("defina seu salario mensal: ")
        fmt.Scanln(&ganha_em_mes)

        if ganha_em_mes < 0 {
            fmt.Println("erro: valor menor que 0")
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

        if ganha_em_mes < 0 {
            fmt.Println("erro: valor menor que 0")
            return
        }

        salario = ganha_em_mes
        fmt.Println("novo salario: ", salario)
        fmt.Println("antigo salario: ", antigo)
    }
}
