// models/transactions.go
package models

import "time"

type Transaction struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Amount      float64   `json:"amount"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Type        string    `json:"type"` // "receita" ou "despesa"
	UserID      uint      `json:"user_id"`
}
