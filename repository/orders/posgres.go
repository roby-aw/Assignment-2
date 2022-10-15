package orders

import (
	"assignment2/business/orders"
	"assignment2/repository"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type PosgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) *PosgresRepository {
	return &PosgresRepository{
		db: db,
	}
}

func (repo *PosgresRepository) CreateOrder(input orders.PostOrder) (id int, err error) {
	order := repository.Order{
		Ordered_at:   input.OrderedAt,
		CustomerName: input.CustomerName,
	}
	err = repo.db.Create(&order).Error
	if err != nil {
		return 0, err
	}

	return int(order.Order_id), nil
}

func (repo *PosgresRepository) CreateItems(input []orders.Items) error {
	for _, item := range input {
		items := repository.Items{
			Order_id:    item.Order_id,
			Quantity:    item.Quantity,
			Item_id:     item.Item_id,
			Item_code:   item.Item_code,
			Description: item.Description,
		}

		err := repo.db.Create(&items).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (repo *PosgresRepository) UpdateItems(input []orders.UpdateItems) error {
	for _, item := range input {
		items := repository.Items{
			Order_id:    item.Order_id,
			Quantity:    item.Quantity,
			Item_id:     item.Item_id,
			Item_code:   item.Item_code,
			Description: item.Description,
		}

		err := repo.db.Model(&repository.Items{}).Where("id = ?", item.LineItemId).Updates(&items).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (repo *PosgresRepository) UpdateOrders(id int, input orders.UpdateOrder) error {
	order := repository.Order{
		Ordered_at:   input.OrderedAt,
		CustomerName: input.CustomerName,
	}

	err := repo.db.Model(&repository.Order{}).Where("order_id = ?", id).Updates(&order).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *PosgresRepository) DeleteOrders(id int) error {
	err := repo.db.Where("order_id = ?", id).Delete(&repository.Order{}, map[string]interface{}{"order_id": id}).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *PosgresRepository) DeleteItems(id int) error {
	err := repo.db.Where("order_id = ?", id).Delete(&repository.Items{}, "order_id = ?", id).Error
	if err != nil {
		return err
	}
	return nil
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
