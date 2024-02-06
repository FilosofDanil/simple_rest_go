package models

type TestUser struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`

	Username string `gorm:"unique;not null" json:"username"`

	TgName string `gorm:"not null" json:"tgName"`

	TgSurname string `gorm:"not null" json:"tgSurname"`

	ChatID uint64 `gorm:"not null" json:"chatId"`
}
