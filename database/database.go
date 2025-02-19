package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := "host=localhost user=postgres password=123 dbname=bvlgedb port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco:", err)
	}
	log.Println("📦 Banco de dados conectado com sucesso!")
}
