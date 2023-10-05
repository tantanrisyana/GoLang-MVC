package main

import (
	"prakerja12/configs"
	"prakerja12/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	loadEnv()
	configs.InitDatabase()
	e := echo.New()
	routes.InitRoutes(e)
	//e.Start(":" + os.Getenv("PORT"))

	e.Start(":" + "8000")
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}
