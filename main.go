package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/gin-contrib/sessions/redis"
	"goblog/platform/authenticator"
	"goblog/platform/database"
	"goblog/platform/router"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load the env vars: %v", err)
	}
	url := os.Getenv("DATABASE_URL")
	err := database.Connect(url)
	if err != nil {
		panic(err)
	}

	auth, err := authenticator.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}
	store, err := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	if err != nil {
		panic(err)
	}

	r := router.New(auth, store)

	log.Print("Server listening on http://localhost:3000/")
	if err := http.ListenAndServe("0.0.0.0:3000", r); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}
}
