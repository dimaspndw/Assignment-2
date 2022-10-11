package models

import (
	"time"
)

type Order struct {
	ID            uint      `gorm:"primaryKey,unique"`
	CustormerName string    `json:"CustomerName" gorm:"not null"`
	OrderedAt     time.Time `json:"OrderedAt" gorm:"not null"`
	Items         []Item    `json:"Item"`
}
