package routers

import (
	"net/http"

	"github.com/fenriz07/twittor/bd"
	"github.com/fenriz07/twittor/models"
)

/*AltaRelacion controlador para realizar un seguimiento*/
func AltaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "id es requerido", http.StatusBadRequest)
		return
	}

	var t models.Relation
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.InsertoRelacion(t)

	if err != nil {
		http.Error(w, "Error al crear la relación", http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "Error al crear la relación", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
