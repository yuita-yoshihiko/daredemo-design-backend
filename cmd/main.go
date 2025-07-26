package main

import (
	"log"
	"net/http"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/yuita-yoshihiko/daredemo-design-backend/adapter/api/router"
	"github.com/yuita-yoshihiko/daredemo-design-backend/config"
)

func main() {
	environment := os.Getenv("ENVIRONMENT")
	if environment == "local" {
		envLoad()
	}
	if err := env.Parse(&config.Conf); err != nil {
		panic(err)
	}

	if err := http.ListenAndServe(":80", router.NewRouter()); err != nil {
		panic(err)
	}
}

func envLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
}
