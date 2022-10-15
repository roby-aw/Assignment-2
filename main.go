package main

import (
	"assignment2/api"
	"assignment2/app/modules"
	"assignment2/config"
	"assignment2/repository"
	"assignment2/utils"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")

	config := config.GetConfig()
	dbCon := utils.NewConnectionDatabase(config)

	defer dbCon.CloseConnection()

	controllers := modules.RegistrationModules(dbCon, config)
	dbCon.Postgres.Migrator().DropTable(&repository.Order{})
	dbCon.Postgres.Migrator().DropTable(&repository.Items{})
	dbCon.Postgres.AutoMigrate(&repository.Order{})
	dbCon.Postgres.AutoMigrate(&repository.Items{})
	//dbCon.Postgres.AutoMigrate(&repository.Customer{})
	//dbCon.Postgres.AutoMigrate(&repository.StockProduct{})
	//dbCon.Postgres.AutoMigrate(&repository.Admin{})

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "API Is Active")
	})
	api.RegistrationPath(e, controllers)

	go func() {
		if port == "" {
			port = "8080"
		}
		address := fmt.Sprintf(":%s", port)
		if err := e.Start(address); err != nil {
			log.Fatal(err)
		}
	}()
	quit := make(chan os.Signal)
	<-quit
}
