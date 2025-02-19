package services

import (
	"github.com/JGMirand4/financial-statistics/database"
	"github.com/JGMirand4/financial-statistics/models"
)

type FinancialStats struct {
	TotalReceitas float64 `json:"total_receitas"`
	TotalDespesas float64 `json:"total_despesas"`
	Saldo         float64 `json:"saldo"`
}

// Função para calcular estatísticas financeiras
func GetFinancialStatistics() (*FinancialStats, error) {
	var totalReceitas float64
	var totalDespesas float64

	// Calcula total de receitas
	err := database.DB.Model(&models.Transaction{}).
		Where("type = ?", "receita").
		Select("SUM(amount)").
		Scan(&totalReceitas)
	if err.Error != nil {
		return nil, err.Error
	}

	// Calcula total de despesas
	err = database.DB.Model(&models.Transaction{}).
		Where("type = ?", "despesa").
		Select("SUM(amount)").
		Scan(&totalDespesas)
	if err.Error != nil {
		return nil, err.Error
	}

	// Retorna os valores
	return &FinancialStats{
		TotalReceitas: totalReceitas,
		TotalDespesas: totalDespesas,
		Saldo:         totalReceitas - totalDespesas,
	}, nil
}
