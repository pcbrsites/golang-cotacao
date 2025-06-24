package servico

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Cotacao struct {
	Bid string `json:"bid"`
}

type CotacaoErro struct {
	StatusCode int    `json:"status_code"`
	Error      string `json:"error"`
	Message    string `json:"message"`
}

func getUrlCotacao() string {
	if os.Getenv("URL_COTACAO") != "" {
		return os.Getenv("URL_COTACAO")
	}
	return "http://localhost:8080/cotacao"
}

func GetCotacaoDolarReal() (*Cotacao, error) {
	log.Println("Iniciando a requisição", getUrlCotacao())
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	defer log.Println("Finalizando a requisição")
	req, err := http.NewRequestWithContext(ctx, "GET", getUrlCotacao(), nil)

	if err != nil {
		return nil, fmt.Errorf("erro na criação da requisição: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Erro na requisição:", err)
		return nil, fmt.Errorf("erro na requisição: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var erro CotacaoErro
		log.Println("Status Code:", resp.StatusCode)
		log.Println("Status:", resp.Status)

		contentType := resp.Header.Get("Content-Type")
		if contentType != "application/json" {
			log.Println("Formato de resposta inesperado:", contentType)
			return nil, fmt.Errorf("erro http: %d, formato: %s", resp.StatusCode, contentType)
		}

		if err := json.NewDecoder(resp.Body).Decode(&erro); err != nil {
			log.Println("Erro ao decodificar resposta de erro:", err)
			return nil, fmt.Errorf("erro ao decodificar resposta de erro: %v status: %d", err, resp.StatusCode)
		}

		log.Println("Erro na requisição:", erro.Error)
		return nil, fmt.Errorf("erro na requisição: %s status: %d", erro.Message, erro.StatusCode)
	}

	var cotacao Cotacao

	if err := json.NewDecoder(resp.Body).Decode(&cotacao); err != nil {
		log.Println("Erro ao decodificar resposta:", err)
		log.Println("Status Code:", resp.StatusCode)
		return nil, fmt.Errorf("erro ao decodificar resposta: %v, status:: %d", err, resp.StatusCode)
	}
	return &cotacao, nil

}

func (cotacao *Cotacao) SalvarCotacaoArquivo() error {
	log.Println("Salvando cotação em arquivo cotacao.txt")
	defer log.Println("Finalizando a gravação do arquivo cotacao.txt")

	file, err := os.Create("cotacao.txt")
	if err != nil {
		log.Println("Erro ao criar arquivo:", err)
		return fmt.Errorf("erro ao criar arquivo: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("Dólar: %s\n", cotacao.Bid))

	if err != nil {
		log.Println("Erro ao escrever no arquivo:", err)
		return fmt.Errorf("erro ao escrever no arquivo: %v", err)
	}
	log.Println("Cotação salva com sucesso")

	return nil
}
