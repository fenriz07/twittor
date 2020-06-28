package bd

import "golang.org/x/crypto/bcrypt"

/*EncriptarPassword se encarga encriptar la contraseña y devolverlo encriptado*/
func EncriptarPassword(pass string) (string, error) {
	costo := 8
	passByte := []byte(pass)
	bytes, err := bcrypt.GenerateFromPassword(passByte, costo)

	return string(bytes), err
}
