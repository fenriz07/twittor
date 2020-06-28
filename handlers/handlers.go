package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/fenriz07/twittor/middleware"
	"github.com/fenriz07/twittor/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* Manejadores seteo mi puerto el handler y pongo a escuchar al servidor*/
func Manejadores() {

	route := mux.NewRouter()

	router.HandleFunc("/registro",middleware.ChequeoBD(routers.Registro)).Methods("POST")



	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(route)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
