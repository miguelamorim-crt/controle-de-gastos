package main

import "fmt"

func compras() {
    var nome_produto string
    var preco_produto float64
    var quantidade_compras int
    var preco_final_compras float64

    fmt.Println("coloque itens na sua lista")

    for {
        fmt.Print("qual o nome do produto: ")
        fmt.Scanln(&nome_produto)
        fmt.Print("qual o valor do produto(0.00 para encerrar): ")
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
