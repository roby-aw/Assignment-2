package api

import (
	"assignment2/api/orders"

	//auth "api-redeem-point/api/middleware"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	OrdersController *orders.Controller
}

func RegistrationPath(e *echo.Echo, controller Controller) {
	e.POST("/orders", controller.OrdersController.CreateOrder)
	e.PUT("/orders/:orderid", controller.OrdersController.UpdateOrders)
	e.DELETE("/orders/:orderid", controller.OrdersController.DeleteOrders)
}
