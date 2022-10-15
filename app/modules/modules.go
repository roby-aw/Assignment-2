package modules

import (
	"assignment2/api"
	ordersApi "assignment2/api/orders"
	ordersBusiness "assignment2/business/orders"
	"assignment2/config"
	ordersRepo "assignment2/repository/orders"
	"assignment2/utils"
)

func RegistrationModules(dbCon *utils.DatabaseConnection, _ *config.AppConfig) api.Controller {

	ordersPermitRepository := ordersRepo.RepositoryFactory(dbCon)
	ordersPermitService := ordersBusiness.NewService(ordersPermitRepository)
	ordersPermitController := ordersApi.NewController(ordersPermitService)

	controller := api.Controller{
		OrdersController: ordersPermitController,
	}
	return controller
}
