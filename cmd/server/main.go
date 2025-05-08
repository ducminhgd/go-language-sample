package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/ducminhgd/go-language-sample/internal/services/translation"
)

func main() {
	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get working directory: %v", err)
	}

	// Create the router
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Create the translation service
	localesDir := filepath.Join(wd, "locales")
	translationService := translation.NewService(localesDir)
	translationHandler := translation.NewHandler(translationService)

	// Routes
	r.Get("/translate", translationHandler.Translate)

	// Health check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Start the server
	port := "8080"
	fmt.Printf("Server starting on port %s...\n", port)
	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
