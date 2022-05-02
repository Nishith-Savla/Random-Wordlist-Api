package main

import (
	"log"
	"os"

	"github.com/Nishith-Savla/Random-Wordlist-Api/app"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("APP_ENV") == "Development" {
		if err := godotenv.Load(); err != nil {
			log.Panic(err)
		}
	}

	app.Start()
}
