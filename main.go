package main

import (
	"log"

	"github.com/Nishith-Savla/Random-Wordlist-Api/app"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	}
	app.Start()
}
