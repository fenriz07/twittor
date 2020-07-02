package routers

import (
	"net/http"

	"github.com/fenriz07/twittor/bd"
	"github.com/fenriz07/twittor/models"
)

/*BajaRelacion controlador de la ruta*/
func BajaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "id es requerido", http.StatusBadRequest)
		return
	}

	var t models.Relation
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.BorroRelacion(t)

	if err != nil {
		http.Error(w, "Error al borrar la relación", http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "Error al borrar la relación", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
