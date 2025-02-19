package controllers

import (
	"net/http"

	"github.com/JGMirand4/financial-statistics/services"

	"github.com/gin-gonic/gin"
)

func GetStatistics(c *gin.Context) {
	stats, err := services.GetFinancialStatistics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao calcular estatísticas"})
		return
	}
	c.JSON(http.StatusOK, stats)
}
