package db

import (
	"context"
	"fmt"
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
	fmt.Println("Iniciando o banco de dados")
	DB, err := gorm.Open(sqlite.Open("cotacao.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Erro ao conectar ao banco de dados:", err)
		return
	}
	Db = DB
	if err := Db.AutoMigrate(&Cotacao{}); err != nil {
		fmt.Println("Erro ao migrar o banco de dados:", err)
		return
	}

}

func NewCotacao(bid string) *Cotacao {
	return &Cotacao{
		Bid: bid,
	}
}

func (cotacao *Cotacao) SalvarCotacao() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*10)

	defer cancel()

	if err := Db.WithContext(ctx).Create(cotacao).Error; err != nil {
		return fmt.Errorf("erro ao salvar cotação: %w", err)
	}

	return nil
}
