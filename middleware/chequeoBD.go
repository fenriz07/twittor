package middleware

import (
	"net/http"

	"github.com/fenriz07/twittor/bd"
)

/*ChequeoBD Middleware que chequea la conexión a la bd*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexión perdida con la base de datos", 500)
			return
		}

		next.ServeHTTP(w, r)
	}
}
