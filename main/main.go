package main

import (
	"go_setup_v1/services"
	"net/http"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=go_setup_v1 port=5432 sslmode=disable"
	services.NewEndpoints(&dsn).Handle()
	println("Server running in 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
