package main

import (
	"chat2/internal/handlers"
	"chat2/internal/models"
	"chat2/internal/services"
	"chat2/internal/web/rest"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	model := models.New()
	fmt.Println("Model layer initialized")

	service := services.New(model)
	fmt.Println("Service layer initialized")

	handler := handlers.New(service)
	fmt.Println("Handler layer initialized")

	r := rest.NewRouter(handler)
	fmt.Println("Routers loaded")

	corsHandler := cors.Default().Handler(r)

	// Start server
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", corsHandler)
}
