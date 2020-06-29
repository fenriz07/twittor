package middleware

import (
	"net/http"

	"github.com/fenriz07/twittor/routers"
)

/*ValidoJWT Middleware para comprobar el Token JWT*/
func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		_, _, _, err := routers.ProcesoToken(r.Header.Get("Authorization"))

		if err != nil {
			http.Error(w, "Token Invalido "+err.Error(), http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	}
}
