package models

import "time"

type Transaction struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Amount      float64   `json:"amount"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Type        string    `json:"type"` // "Income" ou "Loss"
	UserID      uint      `json:"user_id"`
}

// Nome correto da tabela do banco
func (Transaction) TableName() string {
	return "transactions_transaction"
}
