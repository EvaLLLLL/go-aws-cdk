package api

import (
	"fmt"
	"lambda-func/database"
	"lambda-func/types"
)

type ApiHandler struct {
	dbStore database.DynamoDBClient
}

func NewApiHandler(dbStore database.DynamoDBClient) ApiHandler {
	return ApiHandler{
		dbStore: dbStore,
	}
}

func (api ApiHandler) RegisterUserHandler(event types.RegisterUser) error {
	if event.Username == "" || event.Password == "" {
		return fmt.Errorf("request has empty paramters")
	}

	// does a user with this username already exit?

	userExists, err := api.dbStore.DoesUserExist(event.Username)

	if err != nil {
		return fmt.Errorf("there an error checking if user exists - %s", err)
	}

	if userExists {
		return fmt.Errorf("a user with that username already exists")
	}

	// we know that a user does not exist

	insertErr := api.dbStore.InsertUser(event)

	if insertErr != nil {
		return fmt.Errorf("error registering the user - %w", insertErr)
	}

	return nil
}
