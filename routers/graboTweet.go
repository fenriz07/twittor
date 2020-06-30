package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/fenriz07/twittor/bd"
	"github.com/fenriz07/twittor/models"
)

/*GraboTweet controlador de la ruta */
func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet

	err := json.NewDecoder(r.Body).Decode(&mensaje)

	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	if len(mensaje.Mensaje) == 0 {
		http.Error(w, "El mensaje es obligatorio", http.StatusNotAcceptable)
		return
	}

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)

	if err != nil {
		http.Error(w, "No se puede insertar el tweet "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ah logrado insertar el tweet", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
