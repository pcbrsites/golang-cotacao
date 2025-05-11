package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pcbrsites/golang-cotacao/servidor/db"
	httpResponse "github.com/pcbrsites/golang-cotacao/servidor/http"
	"github.com/pcbrsites/golang-cotacao/servidor/servico"
)

func main() {

	db.InitDB()
	http.HandleFunc("/cotacao", handleCotacao)

	fmt.Println("Servidor iniciado na porta http://0.0.0.0:8080/cotacao")
	http.ListenAndServe(":8080", nil)

}

func handleCotacao(w http.ResponseWriter, r *http.Request) {
	cotacao, statusCode, err := servico.GetCotacaoDolarReal()
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		errMsg := httpResponse.ResponseError{
			StatusCode: statusCode,
			Error:      "Erro ao obter cotação",
			Message:    err.Error(),
		}

		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	cotacaoModel := db.NewCotacao(cotacao.Usdbrl.Bid)

	if err := cotacaoModel.SalvarCotacao(); err != nil {

		errMsg := httpResponse.ResponseError{
			StatusCode: http.StatusConflict,
			Error:      "Erro ao salvar cotação",
			Message:    err.Error(),
		}
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	var response httpResponse.ResponseCotacao = httpResponse.ResponseCotacao{
		Bid: cotacao.Usdbrl.Bid,
	}

	json.NewEncoder(w).Encode(response)
}
