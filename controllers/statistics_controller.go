// controllers/statistics_controller.go
package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/JGMirand4/financial-statistics/services"
	"github.com/gin-gonic/gin"
)

// GetStatistics lida com a requisição para obter estatísticas financeiras.
// Parâmetros via query: ?user_id=1&start_date=2020-01-01&end_date=2020-12-31
func GetStatistics(c *gin.Context) {
	// Validação do parâmetro user_id (pode ser definido via middleware de autenticação)
	userIDStr := c.Query("user_id")
	if userIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parâmetro user_id é obrigatório"})
		return
	}
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id inválido"})
		return
	}

	// Parâmetros opcionais para o intervalo de datas
	startDate := c.DefaultQuery("start_date", "1970-01-01")
	endDate := c.DefaultQuery("end_date", time.Now().Format("2006-01-02"))

	stats, err := services.GetFinancialStatistics(c.Request.Context(), uint(userID), startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao calcular estatísticas"})
		return
	}
	c.JSON(http.StatusOK, stats)
}
