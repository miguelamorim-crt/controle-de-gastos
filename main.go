
package main

func main() {
    var opcao int
    for {
        menu()
        fmt.Scanln(&opcao)
        
        separador()
        switch opcao {
        case 0:
            fmt.Println("saindo...")
            creditos()
            return
        case 1:
            adicionarGastos()
        case 2:
            verGastos()
        case 3:
            removerGasto()
        case 4:
            editarGasto()
        case 5:
            criarEspaco()
        case 6:
            compras()
        case 7:
            fmt.Println("salario: ", salario)
        case 8:
            redefinir()
        default:
            fmt.Println("erro: opcao invalida")
        }
        separador()
    }
}
