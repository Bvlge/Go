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
	TotalReceitas          float64 `json:"total_receitas"`
	TotalDespesas          float64 `json:"total_despesas"`
	Saldo                  float64 `json:"saldo"`
	CategoriaMaisFrequente string  `json:"categoria_mais_frequente"`
	TotalTransacoes        int64   `json:"total_transacoes"`
	MediaTransacao         float64 `json:"media_transacao"`
}

// GetFinancialStatistics calcula as estatísticas financeiras para um usuário e intervalo de datas.
// Os parâmetros startDate e endDate devem estar no formato "YYYY-MM-DD".
func GetFinancialStatistics(ctx context.Context, userID uint, startDate, endDate string) (*FinancialStats, error) {
	// Consulta para somar os valores agrupados pelo tipo (Income/Loss).
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
		log.Printf("Erro ao executar consulta de somatório: %v", err)
		return nil, err
	}

	var totalReceitas, totalDespesas float64
	for _, r := range results {
		switch r.Type {
		case "Income":
			totalReceitas = r.Total
		case "Loss":
			totalDespesas = r.Total
		}
	}
	saldo := totalReceitas - totalDespesas

	// Consulta para identificar a categoria que mais aparece
	var topCategory struct {
		Category string `json:"category"`
		Count    int64  `json:"count"`
	}
	err = database.DB.WithContext(ctx).
		Model(&models.Transaction{}).
		Select("category, COUNT(*) as count").
		Where("user_id = ? AND date BETWEEN ? AND ?", userID, startDate, endDate).
		Group("category").
		Order("count DESC").
		Limit(1).
		Scan(&topCategory).Error

	if err != nil {
		log.Printf("Erro ao buscar categoria mais frequente: %v", err)
		return nil, err
	}

	// Consulta para contar o total de transações no período
	var totalTransacoes int64
	err = database.DB.WithContext(ctx).
		Model(&models.Transaction{}).
		Where("user_id = ? AND date BETWEEN ? AND ?", userID, startDate, endDate).
		Count(&totalTransacoes).Error
	if err != nil {
		log.Printf("Erro ao contar transações: %v", err)
		return nil, err
	}

	// Calcula a média das transações (considerando tanto receitas quanto despesas)
	// Aqui assumimos que ambos os valores são somados; se despesas devem ser tratadas como valores negativos, ajuste conforme necessário.
	var mediaTransacao float64
	if totalTransacoes > 0 {
		mediaTransacao = (totalReceitas + totalDespesas) / float64(totalTransacoes)
	}

	log.Printf("Total Receitas: %.2f, Total Despesas: %.2f, Saldo: %.2f, Categoria Top: %s, Total Transações: %d, Média: %.2f",
		totalReceitas, totalDespesas, saldo, topCategory.Category, totalTransacoes, mediaTransacao)

	return &FinancialStats{
		TotalReceitas:          totalReceitas,
		TotalDespesas:          totalDespesas,
		Saldo:                  saldo,
		CategoriaMaisFrequente: topCategory.Category,
		TotalTransacoes:        totalTransacoes,
		MediaTransacao:         mediaTransacao,
	}, nil
}
