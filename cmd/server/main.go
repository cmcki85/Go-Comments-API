package main

import "fmt"

// App - struct that contains things like pointer to DB connections
type App struct{}

// Run - sets up the application
func (app *App) Run() error {
	fmt.Println("Setting up API")
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
