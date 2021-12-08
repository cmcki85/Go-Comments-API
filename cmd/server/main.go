package main

import (
	"fmt"
	"net/http"

	"github.com/cmcki85/RESTurant-Full-Stack/internal/comment"
	"github.com/cmcki85/RESTurant-Full-Stack/internal/database"
	transportHTTP "github.com/cmcki85/RESTurant-Full-Stack/internal/transport/http"
)

// App - struct that contains things like pointer to DB connections
type App struct{}

// Run - sets up the application
func (app *App) Run() error {
	fmt.Println("Setting up API")

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
		fmt.Println("Failed to setup server")
		return err
	}
	return nil
}

func main() {
	fmt.Println("Hello")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting REST api.")
		fmt.Println(err)
	}
}
