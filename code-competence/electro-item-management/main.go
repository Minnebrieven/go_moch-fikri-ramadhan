package main

import (
	"electro-item-management/configs"
	"electro-item-management/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.LoadConfig()
	db, err := configs.ConnectDB()
	if err != nil {
		panic(err)
	}

	err = configs.MigrateAndSeedDB(db)
	if err != nil {
		panic(err)
	}

	e := echo.New()
	routes.NewRoute(db, e)

	e.Logger.Fatal(e.Start(":8000"))
}
