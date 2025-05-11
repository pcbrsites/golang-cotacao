package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Cotacao struct {
	gorm.Model
	Bid string `json:"bid"`
}

var Db *gorm.DB

func InitDB() {
	log.Println("Iniciando o banco de dados")
	DB, err := gorm.Open(sqlite.Open("cotacao.db"), &gorm.Config{})
	if err != nil {
		log.Println("Erro ao conectar ao banco de dados:", err)
		return
	}
	Db = DB
	if err := Db.AutoMigrate(&Cotacao{}); err != nil {
		log.Println("Erro ao migrar o banco de dados:", err)
		return
	}

}

func NewCotacao(bid string) *Cotacao {
	return &Cotacao{
		Bid: bid,
	}
}

func (cotacao *Cotacao) SalvarCotacao() error {
	log.Println("Iniciando a inserção da cotação no banco de dados", cotacao.Bid)
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*10)

	defer cancel()
	defer log.Println("Finalizando a inserção da cotação no banco de dados", cotacao.Bid)

	if err := Db.WithContext(ctx).Create(cotacao).Error; err != nil {
		log.Println("Erro ao salvar cotação:", err)
		return fmt.Errorf("erro ao salvar cotação: %w", err)
	}

	return nil
}
