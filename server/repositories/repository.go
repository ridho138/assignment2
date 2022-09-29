package repositories

import (
	"assignment2/server/models"
)

type OrderRepository interface {
	AddOrder(order *models.Order) error
	GetOrders() (*[]models.DataRes, error)
	AddItems(items *[]models.Item) error
	EditOrder(order *models.Order) error
	EditItems(items *models.Item, id int, itemId int) error
	DeleteOrder(id int) error
	DeleteItem(id int) error
}
