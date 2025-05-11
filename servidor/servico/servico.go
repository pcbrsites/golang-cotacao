package servico

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Cotacao struct {
	Usdbrl Usdbrl `json:"USDBRL"`
}

type Usdbrl struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

func GetCotacaoDolarReal() (*Cotacao, int, error) {
	log.Println("Iniciando a requisição da cotação do dólar")
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*200)
	defer cancel()
	defer log.Println("Finalizando a requisição da cotação do dólar")
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		log.Println("Erro ao criar requisição:", err)
		return nil, http.StatusInternalServerError, fmt.Errorf("erro ao criar requisição: %w", err)
	}

	cliente := &http.Client{}
	res, err := cliente.Do(req)
	if err != nil {
		log.Println("Erro ao fazer requisição:", err)
		return nil, http.StatusRequestTimeout, fmt.Errorf("erro ao fazer requisição: %w", err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Erro ao ler resposta:", err)
		return nil, http.StatusInternalServerError, fmt.Errorf("erro ao ler resposta: %w", err)
	}

	var cotacao Cotacao
	err = json.Unmarshal(body, &cotacao)
	if err != nil {
		log.Println("Erro ao fazer unmarshal:", err)
		return nil, http.StatusInternalServerError, fmt.Errorf("erro ao fazer unmarshal: %w", err)
	}
	if res.StatusCode != http.StatusOK {
		log.Println("Erro ao fazer requisição:", res.Status)
		return nil, res.StatusCode, fmt.Errorf("erro ao fazer requisição: %s", res.Status)
	}
	log.Println("Cotação do dólar:", cotacao.Usdbrl.Bid)
	return &cotacao, http.StatusOK, nil
}
