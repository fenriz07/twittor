package routers

import (
	"encoding/json"
	"net/http"

	"github.com/fenriz07/twittor/bd"
	"github.com/fenriz07/twittor/models"
)

/*ModificarPerfil Controlador de la ruta*/
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Datos Incorrectos "+err.Error(), 400)
		return
	}

	var status bool
	status, err = bd.ModificoRegistro(t, IDUsuario)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el registro, reintente"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se modificaron los datos"+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
