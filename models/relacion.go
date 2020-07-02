package models

/*Ralation modelo de la relacion*/
type Relation struct {
	UsuarioID         string `bson:"usuarioid" json:"user:usuarioId"`
	UsuarioRelacionID string `bson:"usuariorelacionid" json:"user:usuarioRelacionId"`
}
