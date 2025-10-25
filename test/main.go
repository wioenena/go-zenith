package main

import (
	"fmt"

	"github.com/wioenena/go-zenith"
)

func main() {
	app := zenith.NewDefaultApp()
	controllerManager := app.GetControllerManager()
	if err := controllerManager.RegisterControllers(NewMyController()); err != nil {
		fmt.Printf("Error registering controllers: %s\n", err)
		return
	}

	if err := app.Run("localhost:8080"); err != nil {
		panic(err)
	}
}
