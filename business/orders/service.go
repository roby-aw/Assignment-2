package orders

import (
	"github.com/go-playground/validator/v10"
)

type Repository interface {
	CreateOrder(input PostOrder) (id int, err error)
	CreateItems(input []Items) error
	UpdateItems(input []UpdateItems) error
	UpdateOrders(id int, input UpdateOrder) error
	DeleteOrders(id int) error
	DeleteItems(id int) error
}

type Service interface {
	CreateOrder(input PostOrder) error
	CreateItems(input []Items) error
	UpdateItems(input []UpdateItems) error
	UpdateOrders(id int, input UpdateOrder) error
	DeleteOrders(id int) error
}

type service struct {
	repository Repository
	validate   *validator.Validate
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
		validate:   validator.New(),
	}
}

func (s *service) CreateOrder(input PostOrder) (err error) {
	id, err := s.repository.CreateOrder(input)
	if err != nil {
		return err
	}
	var items []Items
	for _, item := range input.Items {
		var tmp Items
		tmp.Order_id = id
		tmp.Item_code = item.Item_code
		tmp.Description = item.Description
		tmp.Quantity = item.Quantity
		tmp.Item_id = item.Item_id
		items = append(items, tmp)
	}
	err = s.repository.CreateItems(items)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) CreateItems(input []Items) error {
	err := s.repository.CreateItems(input)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) UpdateOrders(id int, input UpdateOrder) error {
	err := s.repository.UpdateOrders(id, input)
	if err != nil {
		return err
	}
	err = s.repository.UpdateItems(input.Items)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) UpdateItems(input []UpdateItems) error {
	err := s.repository.UpdateItems(input)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteOrders(id int) error {
	err := s.repository.DeleteOrders(id)
	if err != nil {
		return err
	}
	err = s.repository.DeleteItems(id)
	if err != nil {
		return err
	}
	return nil
}
