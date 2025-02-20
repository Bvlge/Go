// database/database.go
package database

import (
	"log"
	"os"
	"time"

	"github.com/JGMirand4/financial-statistics/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		log.Fatal("DATABASE_DSN não definido")
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	// Configuração do pool de conexões
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Erro ao obter conexão do banco: %v", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Realizar a migração do modelo (apenas para demonstração)
	if err := DB.AutoMigrate(&models.Transaction{}); err != nil {
		log.Fatalf("Erro ao migrar o banco de dados: %v", err)
	}
}
