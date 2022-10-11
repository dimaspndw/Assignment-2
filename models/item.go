package models

type Item struct {
	ID          uint   `gorm:"primaryKey,unique"`
	ItemCode    uint   `json:"ItemCode" gorm:"not null"`
	Description string `json:"Description" gorm:"not null"`
	Quantity    uint   `json:"Quantity" gorm:"not null"`
	OrderID     uint   `gorm:"foreignKey"`
}
