package main

import (
	"github.com/joho/godotenv"
	"os"
	"winton/app/router"
	"winton/config"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		return
	}
	config.InitLog()
}

func main() {
	port := os.Getenv("PORT")

	init := config.Init()
	app := router.Init(init)

	err := app.Run(":" + port)
	if err != nil {
		return
	}
}
