package postgress

import (
	"assignment2/server/models"
	"assignment2/server/repositories"

	"gorm.io/gorm"
)

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) repositories.OrderRepository {
	return &orderRepo{
		db: db,
	}
}

func (o *orderRepo) AddOrder(order *models.Order) error {

	err := o.db.Create(&order).Error
	if err != nil {
		return err
	}
	return nil
}

func (o *orderRepo) AddItems(items *[]models.Item) error {

	err := o.db.Create(&items).Error
	if err != nil {
		return err
	}
	return nil
}

func (o *orderRepo) GetOrders() (*[]models.DataRes, error) {

	var orders []models.Order
	var items []models.Item
	var result []models.DataRes

	err := o.db.Find(&orders).Error
	if err != nil {
		return nil, err
	}

	for _, s := range orders {

		err = o.db.Where("order_id = ?", s.Order_id).Find(&items).Error
		if err != nil {
			return nil, err
		}

		result = append(result, models.DataRes{
			Order_id:      s.Order_id,
			Customer_name: s.Customer_name,
			Ordered_at:    s.Ordered_at,
			Items:         items,
		})
	}

	return &result, nil
}

func (o *orderRepo) EditOrder(order *models.Order) error {

	err := o.db.Save(&order).Error
	if err != nil {
		return err
	}
	return nil
}

func (o *orderRepo) EditItems(items *models.Item, id int, itemId int) error {

	err := o.db.Model(models.Item{}).Where("order_id = ? and item_id = ?", id, itemId).Updates(items).Error
	if err != nil {
		return err
	}
	return nil
}

func (o *orderRepo) DeleteOrder(id int) error {

	err := o.db.Delete(&models.Order{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (o *orderRepo) DeleteItem(id int) error {

	err := o.db.Delete(&models.Item{}, "order_id = ?", id).Error
	if err != nil {
		return err
	}
	return nil
}
