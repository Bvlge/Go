package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/JGMirand4/financial-statistics/controllers"
	"github.com/JGMirand4/financial-statistics/database"
)

func main() {
	// Carrega variáveis de ambiente do arquivo .env, se existir.
	if err := godotenv.Load(); err != nil {
		log.Println("Nenhum arquivo .env encontrado, usando variáveis de ambiente existentes")
	}

	// Verifica se JWT_SECRET está definida; se não, define um valor padrão.
	if os.Getenv("JWT_SECRET") == "" {
		log.Println("JWT_SECRET não definida, utilizando valor padrão '123'")
		os.Setenv("JWT_SECRET", "123")
	}

	// Conectar ao banco de dados.
	database.ConnectDB()

	// Configurar o router do Gin.
	r := gin.Default()

	// Cria um grupo de rotas que requerem autenticação.
	statistics := r.Group("/statistics", controllers.AuthMiddleware())
	{
		statistics.GET("/", controllers.GetStatistics)
		statistics.GET("/category-expenses", controllers.GetMonthlyCategoryExpenses)
	}

	// Iniciar o servidor na porta definida (ou padrão 8080).
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
