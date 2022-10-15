package orders

import (
	"assignment2/business/orders"
	"assignment2/utils"
)

func RepositoryFactory(dbCon *utils.DatabaseConnection) orders.Repository {
	dummyRepo := NewPostgresRepository(dbCon.Postgres)
	return dummyRepo
}
