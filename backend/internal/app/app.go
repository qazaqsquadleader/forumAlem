package app

import (
	"fmt"

	"forum-backend/internal/Log"
	"forum-backend/internal/database"
	"forum-backend/internal/handlers"
)

func Run() error {
	configDB := database.NewConfDB()
	logger, err := Log.CreateLogger() // Setting the logger
	if err != nil {
		return err
	}
	defer Log.CloseLogger(logger)
	if err != nil {
		return err
	}
	db, err := database.InitDB(configDB)
	defer db.Close()
	if err != nil {
		return err
	}
	Log.LogInfo("Successfully Initiated the Data Base")

	if err := database.CreateTables(db); err != nil {
		return err
	}
	Log.LogInfo("Tables have been created")

	apiServer := handlers.NewApiServer(db)

	Log.LogInfo("NewApiServer has been created")
	fmt.Println("Starting the server")
	return apiServer.Start()
}
