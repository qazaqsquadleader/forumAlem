package app

import (
	"forum-backend/internal/database"
	"forum-backend/internal/handler"
)

func Run() error {
	configDB := database.NewConfDB()
	db, err := database.InitDB(configDB)
	if err != nil {
		return err
	}
	if err := database.CreateTables(db); err != nil {
		return err
	}
	apiServer := handler.NewApiServer(db)

	return apiServer.Start()
}
