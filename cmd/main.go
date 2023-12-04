package main

import (
	"flag"
	"log"
)

func main() {
	token := token()
}

func token() string {
	token := flag.String("token-bot-token", "", "token for telegram bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("No tokens?:()")
	}

	return *token
}
