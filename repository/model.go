package repository

import (
	"time"
)

type Order struct {
	Order_id     uint      `gorm:"primaryKey;autoIncrement:true"`
	CustomerName string    `gorm:"type:varchar(255)" json:"customerName"`
	Ordered_at   time.Time `json:"orderedAt"`
	// Items        Items
}

type Items struct {
	ID          uint   `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Order_id    int    `json:"order_id"`
	Quantity    int    `json:"quantity"`
	Item_id     int    `json:"item_id"`
	Item_code   string `gorm:"type:varchar(255)" json:"item_code"`
	Description string `gorm:"type:varchar(255)" json:"description"`
}
