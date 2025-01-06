package entity

import (
	"time"
)

type Stock struct {
	ID          int       `gorm:"primaryKey;autoIncrement"`
	OfferID     string    `gorm:"size:30;index"`
	AccountName string    `gorm:"not null"`
	LastStock   int       `gorm:"not null"` // LastStock последний остаток на площадке
	NewStock    int       `gorm:"not null"` // NewStock самый актуальный остаток
	LastUpdated time.Time `gorm:"not null"` // LastUpdated дата последнего обновления на площадке
	NewUpdated  time.Time `gorm:"not null"` // NewUpdated дата самого актуального остатка
}
