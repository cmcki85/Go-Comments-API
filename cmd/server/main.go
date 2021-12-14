package main

import (
	"net/http"

	"github.com/cmcki85/RESTurant-Full-Stack/internal/comment"
	"github.com/cmcki85/RESTurant-Full-Stack/internal/database"
	transportHTTP "github.com/cmcki85/RESTurant-Full-Stack/internal/transport/http"

	log "github.com/sirupsen/logrus"
)

// App - Contains the application information
type App struct{
	Name string
	Version string
}

// Run - sets up the application
func (app *App) Run() error {
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields{
			"AppName": app.Name,
			"AppVersion": app.Version,
		}).Info("Setting up application")

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}

	err = database.MigrateDB(db)
	if err != nil {
		return err
	}

	commentService := comment.NewService(db)

	handler := transportHTTP.NewHandler(commentService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		log.Error("Failed to setup server")
		return err
	}
	return nil
}

func main() {
	app := App{
		Name: "Commenting service",
		Version: "1.0.0",
	}
	if err := app.Run(); err != nil {
		log.Error("Error starting REST api.")
		log.Fatal(err)
	}
}
