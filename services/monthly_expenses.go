package services

import (
	"context"
	"log"

	"github.com/JGMirand4/financial-statistics/database"
	"github.com/JGMirand4/financial-statistics/models"
)

// CategoryExpense representa a média de despesas para uma categoria em um determinado mês.
type CategoryExpense struct {
	Category     string  `json:"category"`
	YearMonth    string  `json:"year_month"` // Formato YYYY-MM
	AvgExpense   float64 `json:"avg_expense"`
	TotalExpense float64 `json:"total_expense"`
	Count        int64   `json:"count"`
}

// GetMonthlyCategoryExpenses calcula a média mensal de despesas por categoria para um usuário num intervalo de datas.
func GetMonthlyCategoryExpenses(ctx context.Context, userID uint, startDate, endDate string) ([]CategoryExpense, error) {
	var results []CategoryExpense

	// A query abaixo utiliza a função DATE_TRUNC para agrupar as transações pelo mês.
	err := database.DB.WithContext(ctx).
		Model(&models.Transaction{}).
		Select(`category, 
		        TO_CHAR(date, 'YYYY-MM') as year_month, 
		        AVG(amount) as avg_expense,
		        SUM(amount) as total_expense,
		        COUNT(*) as count`).
		Where("user_id = ? AND type IN ? AND date BETWEEN ? AND ?", userID, []string{"despesa", "expense"}, startDate, endDate).
		Group("category, year_month").
		Order("year_month ASC").
		Scan(&results).Error

	if err != nil {
		log.Printf("Erro ao executar consulta de despesas por categoria: %v", err)
		return nil, err
	}

	log.Printf("Resultados da consulta de despesas por categoria: %+v", results)
	return results, nil
}
