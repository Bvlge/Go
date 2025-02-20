// main.go
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
	// Carrega variáveis de ambiente do arquivo .env (caso exista)
	if err := godotenv.Load(); err != nil {
		log.Println("Nenhum arquivo .env encontrado, usando variáveis de ambiente existentes")
	}

	// Conectar ao banco de dados
	database.ConnectDB()

	// Configurar o router do Gin
	r := gin.Default()

	// (Opcional) Adicionar middleware de autenticação
	// r.Use(controllers.AuthMiddleware())

	// Definir as rotas
	r.GET("/statistics", controllers.GetStatistics)

	// Iniciar o servidor na porta definida (ou padrão 8080)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
