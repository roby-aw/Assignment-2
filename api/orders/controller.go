package orders

import (
	ordersBussiness "assignment2/business/orders"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service ordersBussiness.Service
}

func NewController(service ordersBussiness.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (Controller *Controller) CreateOrder(e echo.Context) error {
	var payload ordersBussiness.PostOrder
	e.Bind(&payload)
	err := Controller.service.CreateOrder(payload)
	if err != nil {
		return err
	}
	return e.JSON(200, map[string]interface{}{
		"message": "success",
	})
}

func (Controller *Controller) UpdateOrders(c echo.Context) error {
	var payload ordersBussiness.UpdateOrder
	id := c.Param("orderid")
	orderid, _ := strconv.Atoi(id)
	c.Bind(&payload)
	err := Controller.service.UpdateOrders(orderid, payload)
	if err != nil {
		return err
	}
	return c.JSON(200, map[string]interface{}{
		"message": "success",
	})
}

func (Controller *Controller) DeleteOrders(c echo.Context) error {
	id := c.Param("orderid")
	orderid, _ := strconv.Atoi(id)
	err := Controller.service.DeleteOrders(orderid)
	if err != nil {
		return err
	}
	return c.JSON(200, map[string]interface{}{
		"message": "success",
	})
}
