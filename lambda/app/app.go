package app

import (
	"lambda-func/api"
	"lambda-func/database"
)

type App struct {
	ApiHandler api.ApiHandler
}

func NewApp() App {
	// we actually initialize our DB store
	// gets passed down into the api handler
	db := database.NewDynamoDbClient()
	apiHandler := api.NewApiHandler(db)

	return App{
		ApiHandler: apiHandler,
	}
}
