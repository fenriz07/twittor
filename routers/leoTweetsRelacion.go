package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/fenriz07/twittor/bd"
)

/*LeoTweetsSeguidoresRelacion controlador de la ruta*/
func LeoTweetsSeguidoresRelacion(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "pagina es requerido", http.StatusBadRequest)
		return
	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	respuesta, status := bd.LeoTweetsSeguidores(IDUsuario, pagina)

	if status == false {
		http.Error(w, "Error consultando los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(respuesta)

}
