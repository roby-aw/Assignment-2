package orders

import "time"

type Order struct {
	Order_id     int    `json:"order_id"`
	CustomerName string `gorm:"type:varchar(255)" json:"customerName"`
}

type Items struct {
	Order_id    int    `json:"order_id"`
	Quantity    int    `json:"quantity"`
	Item_id     int    `json:"item_id"`
	Item_code   string `gorm:"type:varchar(255)" json:"itemCode"`
	Description string `gorm:"type:varchar(255)" json:"description"`
}

type UpdateItems struct {
	LineItemId  int    `json:"lineItemId"`
	Order_id    int    `json:"order_id"`
	Quantity    int    `json:"quantity"`
	Item_id     int    `json:"item_id"`
	Item_code   string `gorm:"type:varchar(255)" json:"itemCode"`
	Description string `gorm:"type:varchar(255)" json:"description"`
}

type PostOrder struct {
	OrderedAt    time.Time `json:"orderedAt"`
	CustomerName string    `gorm:"type:varchar(255)" json:"customerName"`
	Items        []Items   `json:"items"`
}

type UpdateOrder struct {
	Order_id     int           `json:"order_id"`
	OrderedAt    time.Time     `json:"orderedAt"`
	CustomerName string        `gorm:"type:varchar(255)" json:"customerName"`
	Items        []UpdateItems `json:"items"`
}
