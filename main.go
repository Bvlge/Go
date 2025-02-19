package main

import (
	"github.com/JGMirand4/financial-statistics/controllers"
	"github.com/JGMirand4/financial-statistics/database"
	"github.com/gin-gonic/gin"
)

func main() {
	// Conectar ao banco de dados
	database.ConnectDB()

	r := gin.Default()

	// Definir as rotas
	r.GET("/statistics", controllers.GetStatistics)

	// Iniciar servidor
	r.Run(":8080")
}
