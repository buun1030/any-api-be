package main

import (
	"log"
	"net/http"

	"any-api/internal/config"
	"any-api/internal/handlers"
	"any-api/internal/repository"
	"any-api/internal/services"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize repositories
	messageRepo := repository.NewInMemoryMessageRepository()
	itemRepo, err := repository.NewPostgresItemRepository(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to initialize item repository: %v", err)
	}

	// Initialize services
	helloService := services.NewHelloService(messageRepo)
	itemService := services.NewItemService(itemRepo)

	// Initialize handlers
	helloHandler := &handlers.HelloHandler{Service: helloService}
	itemHandler := &handlers.ItemHandler{Service: itemService}

	r := mux.NewRouter()
	r.Handle("/hello", helloHandler).Methods("GET")
	r.HandleFunc("/items", itemHandler.CreateItem).Methods("POST")

	c := cors.Default()
	handler := c.Handler(r)

	log.Printf("Starting server on %s\n", cfg.Port)
	err = http.ListenAndServe(cfg.Port, handler)
	if err != nil {
		log.Fatal(err)
	}
}
