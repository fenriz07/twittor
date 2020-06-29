package models

/*RespuestaLogin tiene el token que se le va a devolver en el login*/
type RespuestaLogin struct{
	Token string `json:"token,omiempty"`
}