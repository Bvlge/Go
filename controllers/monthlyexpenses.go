package controllers

import (
	"net/http"
	"time"

	"github.com/JGMirand4/financial-statistics/services"
	"github.com/gin-gonic/gin"
)

func GetMonthlyCategoryExpenses(c *gin.Context) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	userID, ok := userIDInterface.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao identificar usuário"})
		return
	}

	// Parâmetros de data
	startDateStr := c.DefaultQuery("start_date", "2023-01-01")
	endDateStr := c.DefaultQuery("end_date", time.Now().Format("2006-01-02"))

	if _, err := time.Parse("2006-01-02", startDateStr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de start_date inválido. Use YYYY-MM-DD"})
		return
	}

	if _, err := time.Parse("2006-01-02", endDateStr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de end_date inválido. Use YYYY-MM-DD"})
		return
	}

	results, err := services.GetMonthlyCategoryExpenses(c.Request.Context(), userID, startDateStr, endDateStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao calcular estatísticas por categoria"})
		return
	}

	c.JSON(http.StatusOK, results)
}
