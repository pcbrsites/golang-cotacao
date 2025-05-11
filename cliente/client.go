package main

import (
	"fmt"

	"github.com/pcbrsites/golang-cotacao/cliente/servico"
)

func main() {

	cotacao, err := servico.GetCotacaoDolarReal()
	if err != nil {
		fmt.Println("Erro ao obter cotação:", err)
		return
	}
	fmt.Println("Cotação do dolar:", cotacao.Bid)

	err = cotacao.SalvarCotacaoArquivo()
	if err != nil {
		fmt.Println("Erro ao salvar cotação em arquivo:", err)
		return
	}
}
