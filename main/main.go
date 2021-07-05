package main

import (
	"go_setup_v1/services/auth"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=go_setup_v1 port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		panic(err)
	}
	auth.NewEndpoints(db).Handle()
	println("Server running in 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
