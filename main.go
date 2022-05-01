package main

import (
	"log"

	"github.com/Nishith-Savla/Random-Wordlist-Api/app"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Panic(err)
	}
	app.Start()
}
