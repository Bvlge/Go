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
	var results []struct {
		Type  string  `json:"type"`
		Total float64 `json:"total"`
	}

	log.Printf("Consultando estatísticas para user_id=%d entre %s e %s", userID, startDate, endDate)

	err := database.DB.WithContext(ctx).
		Model(&models.Transaction{}).
		Select("type, COALESCE(SUM(amount), 0) as total").
		Where("user_id = ? AND date BETWEEN ? AND ?", userID, startDate, endDate).
		Group("type").
		Scan(&results).Error

	if err != nil {
		log.Println("Erro ao executar consulta SQL:", err)
		return nil, err
	}

	log.Printf("Resultados da consulta: %+v", results)

	var totalReceitas, totalDespesas float64
	for _, r := range results {
		if r.Type == "receita" || r.Type == "income" {
			totalReceitas = r.Total
		} else if r.Type == "despesa" || r.Type == "expense" {
			totalDespesas = r.Total
		}
	}

	log.Printf("Total Receitas: %.2f, Total Despesas: %.2f, Saldo: %.2f", totalReceitas, totalDespesas, totalReceitas-totalDespesas)

	return &FinancialStats{
		TotalReceitas: totalReceitas,
		TotalDespesas: totalDespesas,
		Saldo:         totalReceitas - totalDespesas,
	}, nil
}
