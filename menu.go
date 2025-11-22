package main

import "fmt"

func menu() {
    fmt.Println("====  ControleMirim =====")
    fmt.Println("1 - adicionar gasto")
    fmt.Println("2 - ver gastos")
    fmt.Println("3 - remover gasto")
    fmt.Println("4 - editar gasto")
    fmt.Println("5 - lista de compras")
    fmt.Println("6 - adicionar espaço")
    fmt.Println("7 - ver salario")
    fmt.Println("8 - definir ou redefinir salario mensal")
    fmt.Println("0 - sair")
    fmt.Println("=========================")
    fmt.Print("opcao escolhida: ")
}

func separador() {
    fmt.Println("================================")
}

func creditos() {
    fmt.Println("================================")
    fmt.Println("Autor: Miguel A.")
    fmt.Println("Versão: 1.1.0")
    fmt.Println("Ano: 2025")
    fmt.Println("espaços inspirados em Perplexity AI")
    fmt.Println("primeira versao desenvolvida em python")
    fmt.Println("================================")
}
