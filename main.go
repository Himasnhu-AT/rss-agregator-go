package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable not set")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler((cors.Options{
		AllowedOrigins:   []string{"https://", "http:"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})))
	serve := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	fmt.Printf("Listening on port %v\n", port)

	err := serve.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
