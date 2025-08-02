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

	// Middleware to handle X-Forwarded-Proto header
	proxyMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if xfp := r.Header.Get("X-Forwarded-Proto"); xfp == "https" {
				r.URL.Scheme = "https"
			}
			next.ServeHTTP(w, r)
		})
	}

	c := cors.Default()
	handler := c.Handler(proxyMiddleware(r))

	log.Printf("Starting server on %s\n", cfg.Port)
	err = http.ListenAndServe(cfg.Port, handler)
	if err != nil {
		log.Fatal(err)
	}
}
