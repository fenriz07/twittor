package bd

import (
	"github.com/fenriz07/twittor/models"
	"golang.org/x/crypto/bcrypt"
)

/*IntentoLogin verifica si la contraseña que envio  el usuario es correcta*/
func IntentoLogin(email string, password string) (models.Usuario, bool) {
	usu, encontrado, _ := ChequeoYaExisteUsuario(email)

	if encontrado == false {
		return usu, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	if err != nil {
		return usu, false
	}

	return usu, true
}
