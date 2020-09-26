package main

import (
	"go_setup_v1/services/auth"
	"net/http"
)

func main() {
	auth.NewEndpoints().Handle()
	println("Server running in 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
