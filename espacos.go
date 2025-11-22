package main

import "fmt"

func criarEspaco() {
    var nome_espaco string
    var quantidade_espaco int

    fmt.Print("qual o nome: ")
    fmt.Scanln(&nome_espaco)

    fmt.Printf("quantos gastos quer adicionar no espaco %s: ", nome_espaco)
    fmt.Scanln(&quantidade_espaco)

    if quantidade_espaco < 0 {
        fmt.Println("erro: quantidade menor que zero")
        return
    }

    if quantidade_espaco > len(gastos) {
        fmt.Println("erro: quantidade maior que quantidade de gastos registrados")
        return
    }

    if len(gastos) == 0 {
        fmt.Println("erro: nao existe gastos para serem adicionados ao espaco")
        return
    }

    for i := 0; i < len(gastos); i++ {
        fmt.Printf("%d. nome: %s | gasto: R$ %.2f\n", i+1, gastos[i], valores[i])
    }

    for c := 1; c <= quantidade_espaco; c++ {
        fmt.Printf("Digite o número do gasto para adicionar ao espaço %s: ", nome_espaco)
        var indiceGasto int
        fmt.Scanln(&indiceGasto)
        indiceGasto--

        if indiceGasto < 0 || indiceGasto >= len(gastos) {
            fmt.Println("Índice inválido! Tente novamente.")
            c--
            continue
        }

        gastos_espaco = append(gastos_espaco, gastos[indiceGasto])
        fmt.Printf("Gasto \"%s\" adicionado ao espaço \"%s\"!\n", gastos[indiceGasto], nome_espaco)
    }
    fmt.Printf("Espaço \"%s\" criado com os gastos: %v\n", nome_espaco, gastos_espaco)
}
