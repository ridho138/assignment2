package models

import "time"

type Item struct {
	Item_id     int `gorm:"primaryKey"`
	Item_code   string
	Description string
	Quantity    int
	Order_id    int
}

type Order struct {
	Order_id      int `gorm:"primaryKey"`
	Customer_name string
	Ordered_at    time.Time
}

type ItemReq struct {
	Item_id     int
	Item_code   string
	Description string
	Quantity    int
}

type DataReq struct {
	Customer_name string
	Items         []ItemReq
}

type DataRes struct {
	Order_id      int
	Customer_name string
	Ordered_at    time.Time
	Items         []Item
}
