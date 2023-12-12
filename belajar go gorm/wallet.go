package belajar_go_gorm

import "time"

type Wallet struct {
	ID        string    `gorm:"primary_key;column:id"`
	UserId    string    `gorm:"column:user_id"`
	Balance   int64     `gorm:"column:balance"`
	CreatedAt time.Time `gorm:"colum:created_at"`
	UpdatedAt time.Time `gorm:"colum:updated_at"`
}

func (w Wallet) TableName() string {
	return "wallets"
}
