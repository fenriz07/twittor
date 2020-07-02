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

/*Manejadores seteo mi puerto el handler y pongo a escuchar al servidor*/
func Manejadores() {

	router := mux.NewRouter()

	router.HandleFunc("/registro", middleware.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middleware.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middleware.ChequeoBD(middleware.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middleware.ChequeoBD(middleware.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middleware.ChequeoBD(middleware.ValidoJWT(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/leoTweets", middleware.ChequeoBD(middleware.ValidoJWT(routers.LeoTweets))).Methods("GET")
	router.HandleFunc("/eliminarTweet", middleware.ChequeoBD(middleware.ValidoJWT(routers.EliminarTweet))).Methods("DELETE")

	router.HandleFunc("/subirAvatar", middleware.ChequeoBD(middleware.ValidoJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/subirBanner", middleware.ChequeoBD(middleware.ValidoJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/obtenerAvatar", middleware.ChequeoBD(routers.ObtenerAvatar)).Methods("GET")
	router.HandleFunc("/obtenerBanner", middleware.ChequeoBD(routers.ObtenerBanner)).Methods("GET")

	router.HandleFunc("/altaRelacion", middleware.ChequeoBD(middleware.ValidoJWT(routers.AltaRelacion))).Methods("GET")
	router.HandleFunc("/bajaRelacion", middleware.ChequeoBD(middleware.ValidoJWT(routers.BajaRelacion))).Methods("GET")

	router.HandleFunc("/consultaRelacion", middleware.ChequeoBD(middleware.ValidoJWT(routers.ConsultaRelacion))).Methods("GET")
	router.HandleFunc("/listaUsuarios", middleware.ChequeoBD(middleware.ValidoJWT(routers.ListaUsuarios))).Methods("GET")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
