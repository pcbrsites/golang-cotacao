package main

import (
	"log"

	"github.com/pcbrsites/golang-cotacao/cliente/servico"
)

func main() {

	cotacao, err := servico.GetCotacaoDolarReal()
	if err != nil {
		log.Println("Erro ao obter cotação:", err)
		return
	}
	log.Println("Cotação do dolar:", cotacao.Bid)

	err = cotacao.SalvarCotacaoArquivo()
	if err != nil {
		log.Println("Erro ao salvar cotação em arquivo:", err)
		return
	}
}
