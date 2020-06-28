package main

import (
	"log"

	"github.com/fenriz07/twittor/bd"
	"github.com/fenriz07/twittor/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexipn a la bd")
		return
	}

	handlers.Manejadores()
}
