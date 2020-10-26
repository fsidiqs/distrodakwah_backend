package app

import (
	"fmt"
	"os"
	"strconv"

	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/database"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/router"
	product "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/router/product"
)

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

	port := 3001
	env := os.Getenv("ENV")
	if env == "STAGING" || env == "PRODUCTION" {
		port, _ = strconv.Atoi(os.Getenv("PORT"))
	}
	fmt.Printf("App is running on port: %v", port)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
