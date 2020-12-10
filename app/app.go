package app

import (
	"fmt"
	"os"
	"strconv"

	"distrodakwah_backend/app/database"
	"distrodakwah_backend/app/router"
	product "distrodakwah_backend/app/router/product"
	"distrodakwah_backend/config"
)

func init() {
	config.Load()
}
func Run() {
	fmt.Println("Connecting to db")
	err := database.Connect()
	if err != nil {
		fmt.Println("Failed to conenct db")
	}
	fmt.Println("Database Connected")
	// init api
	product.Init()

	e := router.Init()

	fmt.Println("the port is:", os.Getenv("Port"))
	env := os.Getenv("ENV")
	// port, _ := strconv.Atoi(os.Getenv("PORT"))
	port := 3001
	if env == "STAGING" || env == "PRODUCTION" {
		port, _ = strconv.Atoi(os.Getenv("PORT"))
	}
	fmt.Println("App is running on", port)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
