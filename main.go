package main

import (
	"finance-tracker/controller"
	"finance-tracker/model"
	"fmt"
	"os"

	"github.com/google/uuid"
)

func main() {
	os.Setenv("API_SECRET", uuid.NewString())
	ds := fmt.Sprintf("port=%d host=%s user=%s password=%s dbname=%s sslmode=disable", 5432, "db", "postgres", "postgres", "postgres")
	db, err := model.NewDB(ds)
	if err != nil {
		fmt.Printf("Failed to initialize database: %v\n", err)
		os.Exit(1)
	}
	server := controller.Server{
		DB: db,
	}
	server.Initialize()
}
