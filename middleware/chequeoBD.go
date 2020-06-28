package middleware

import (
	"net/http"

	"github.com/fenriz07/twittor/bd"
)

func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http
		}
	}
}
