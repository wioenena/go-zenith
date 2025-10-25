package main

import (
	"github.com/wioenena/go-zenith"
)

func main() {
	application := zenith.DefaultApp()
	if err := application.Run("localhost:8080"); err != nil {
		panic(err)
	}
}
