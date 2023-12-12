package belajar_go_gorm

import "time"

type User struct {
	ID        string    `gorm:"primary_key;column:id;<-:create"`
	Password  string    `gorm:"colum:password"`
	Name      string    `gorm:"colum:name"`
	CreatedAt time.Time `gorm:"colum:created_at;autoCreateTime;<-:create"`
	UpdatedAt time.Time `gorm:"colum:updated_at;autoCreateTime;autoUpdateTime"`
	Wallet    Wallet    `gorm:"foreignKey:user_id;references:id"`
}

func (u User) TableName() string {
	return "users"
}
