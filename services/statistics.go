// services/statistics.go
package services

import (
	"context"
	"log"

	"github.com/JGMirand4/financial-statistics/database"
	"github.com/JGMirand4/financial-statistics/models"
)

// FinancialStats armazena as estatísticas financeiras calculadas.
type FinancialStats struct {
	TotalReceitas float64 `json:"total_receitas"`
	TotalDespesas float64 `json:"total_despesas"`
	Saldo         float64 `json:"saldo"`
}

// GetFinancialStatistics calcula as estatísticas financeiras para um usuário e intervalo de datas.
// Os parâmetros startDate e endDate devem estar no formato "YYYY-MM-DD".
func GetFinancialStatistics(ctx context.Context, userID uint, startDate, endDate string) (*FinancialStats, error) {
	var totalReceitas float64
	var totalDespesas float64

	// Consulta para somar as receitas filtrando por usuário e intervalo de datas
	err := database.DB.WithContext(ctx).
		Model(&models.Transaction{}).
		Where("type = ? AND user_id = ? AND date BETWEEN ? AND ?", "receita", userID, startDate, endDate).
		Select("COALESCE(SUM(amount), 0)").
		Row().
		Scan(&totalReceitas)
	if err != nil {
		log.Printf("Erro ao calcular receitas: %v", err)
		return nil, err
	}

	// Consulta para somar as despesas filtrando por usuário e intervalo de datas
	err = database.DB.WithContext(ctx).
		Model(&models.Transaction{}).
		Where("type = ? AND user_id = ? AND date BETWEEN ? AND ?", "despesa", userID, startDate, endDate).
		Select("COALESCE(SUM(amount), 0)").
		Row().
		Scan(&totalDespesas)
	if err != nil {
		log.Printf("Erro ao calcular despesas: %v", err)
		return nil, err
	}

	return &FinancialStats{
		TotalReceitas: totalReceitas,
		TotalDespesas: totalDespesas,
		Saldo:         totalReceitas - totalDespesas,
	}, nil
}
